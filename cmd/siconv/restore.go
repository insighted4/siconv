package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"runtime"

	"github.com/go-pg/pg"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/server"
	"github.com/insighted4/siconv/storage"
	"github.com/insighted4/siconv/storage/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func commandRestore() *cobra.Command {
	var (
		file        string
		databaseURL string
		truncate    bool
		logLevel    string
		logFormat   string
	)
	cmd := cobra.Command{
		Use:     "restore",
		Short:   "Restore database from an archive file",
		Example: "siconv restore",
		Run: func(cmd *cobra.Command, args []string) {
			parsedURL, err := url.Parse(databaseURL)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			filename, err := filepath.Abs(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid file: %s", filename)
				os.Exit(2)
			}

			ext := filepath.Ext(filename)
			if ext != ".zip" && ext != ".csv" {
				fmt.Fprintf(os.Stderr, "%s is not a ZIP or CSV file", filename)
				os.Exit(2)
			}

			pwd, _ := parsedURL.User.Password()
			pgOptions := &pg.Options{
				Addr:     parsedURL.Host,
				Database: strings.Replace(parsedURL.RequestURI(), "/", "", 1),
				User:     parsedURL.User.Username(),
				Password: pwd,
			}

			logger, err := server.NewLogger(logLevel, logFormat)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			start := time.Now()
			if err := restore(filename, truncate, pgOptions, logger); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			logger.Infof("Total: %.2fmin", time.Since(start).Minutes())
		},
	}

	cmd.Flags().StringVar(&file, "file", "siconv.zip", "ZIP or CSV file")
	cmd.Flags().StringVar(&databaseURL, "database-url", "postgres://localhost:5432/siconv", "Database connection string")
	cmd.Flags().BoolVar(&truncate, "truncate", false, "Truncate table")
	cmd.Flags().StringVar(&logLevel, "log-level", "info", "Logger level")
	cmd.Flags().StringVar(&logFormat, "log-format", "text", "Logger format")

	return &cmd
}

func restore(filename string, truncate bool, pgOptions *pg.Options, logger logrus.FieldLogger) error {
	logger.Infof("Restoring database from %s", filename)

	storage := postgres.New(pgOptions, logger)

	switch filepath.Ext(filename) {
	case ".csv":
		if err := restoreCSV(filename, truncate, storage, logger); err != nil {
			return err
		}
	case ".zip":
		if err := restoreZIP(filename, truncate, storage, logger); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unrecognized file extension")
	}

	return nil
}

func restoreCSV(filename string, truncate bool, storage storage.Service, logger logrus.FieldLogger) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return insert(file, filepath.Base(filename), truncate, storage, logger)
}

func restoreZIP(filename string, truncate bool, storage storage.Service, logger logrus.FieldLogger) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		if err := insertZIP(file, truncate, storage, logger); err != nil {
			return err
		}
	}

	return nil
}

func insertZIP(file *zip.File, truncate bool, storage storage.Service, logger logrus.FieldLogger) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	filename := file.FileInfo().Name()
	logger.Infof("Processing %s", filename)
	if err := insert(fileReader, filename, truncate, storage, logger); err != nil {
		return err
	}

	return nil
}

func insert(fileReader io.Reader, filename string, truncate bool, storage storage.Service, logger logrus.FieldLogger) error {
	start := time.Now()
	converter := NewConverter(filename)
	if converter == nil {
		logger.Warnf("unrecognized filename: %s", filename)
		return nil
	}

	if truncate {
		tablename := switchFileToTable(filename)
		if tablename == "" {
			logger.Warnf("unrecognized table name: %s", filename)
		}
		if err := storage.Truncate(tablename); err != nil {
			return err
		}
	}

	csvReader := csv.NewReader(fileReader)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	csvReader.TrimLeadingSpace = true

	headers, err := csvReader.Read()
	if err != nil {
		return err
	}

	line := 2
	limit := 1000
	wg := NewWaitGroup(runtime.NumCPU())
	for {
		rows, err := readRows(csvReader, headers, limit)
		if err != nil {
			return err
		}

		if len(rows) == 0 {
			break
		}

		wg.Add()
		go func() {
			defer wg.Done()
			var models []schema.Model
			for _, row := range rows {
				model := converter(row)
				if model != nil {
					model.SetID(line)
					models = append(models, model)
				}

				if (line % limit) == 0 {
					logger.Debugf("Processed %s: %d lines", filename, line)
				}

				line++
			}

			if err := storage.BulkInsert(&models); err != nil {
				logger.Errorf("unable to process %s:%d: %v", filename, line, err)
			}

		}()

		wg.Wait()
	}

	logger.Infof("Processed %s: %d lines took %.2fmin", filename, line, time.Since(start).Minutes())
	return nil
}

func readRows(reader *csv.Reader, headers []string, limit int) ([]map[string]string, error) {
	models := []map[string]string{}
	total := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		row := map[string]string{}
		for index, header := range headers {
			h := strings.TrimSpace(strings.ToLower(header))
			h = strings.Trim(h, "\xef\xbb\xbf") // Remove infamous BOM

			v := strings.TrimSpace(record[index])
			row[h] = v
		}

		models = append(models, row)
		total++

		if total >= limit {
			break
		}
	}

	return models, nil
}
