package main

import (
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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func commandDatabase() *cobra.Command {
	var (
		file        string
		databaseURL string
		logLevel    string
		logFormat   string
	)
	cmd := cobra.Command{
		Use:     "restore",
		Short:   "Restore database from an archive file",
		Example: "database restore",
		Run: func(cmd *cobra.Command, args []string) {
			u, err := url.Parse(viper.GetString("database_url"))
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

			pwd, _ := u.User.Password()
			pgOptions := &pg.Options{
				Addr:     u.Host,
				Database: strings.Replace(u.RequestURI(), "/", "", 1),
				User:     u.User.Username(),
				Password: pwd,
			}

			logger, err := server.NewLogger(viper.GetString("log_level"), viper.GetString("log_format"))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			start := time.Now()
			if err := restore(filename, pgOptions, logger); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			logger.Infof("Total: %.2fmin", time.Since(start).Minutes())
		},
	}

	viper.SetEnvPrefix("siconv")
	viper.AutomaticEnv()

	cmd.Flags().StringVar(&file, "file", "siconv.zip", "ZIP or CSV file")

	cmd.Flags().StringVar(&databaseURL, "database-url", "postgres://localhost:5432/siconv", "Database connection string")
	viper.BindPFlag("database_url", cmd.Flags().Lookup("database-url"))

	cmd.Flags().StringVar(&logLevel, "log-level", "info", "Logger level")
	viper.BindPFlag("log_level", cmd.Flags().Lookup("log-level"))

	cmd.Flags().StringVar(&logFormat, "log-format", "text", "Logger format")
	viper.BindPFlag("log_format", cmd.Flags().Lookup("log-format"))

	return &cmd
}

func restore(filename string, pgOptions *pg.Options, logger logrus.FieldLogger) error {
	logger.Infof("Updating database %s", filename)

	storage := NewPostgres(pgOptions, logger)

	switch filepath.Ext(filename) {
	case ".csv":
		if err := restoreCSV(filename, storage, logger); err != nil {
			return err
		}
	case ".zip":
		if err := restoreZIP(filename, storage, logger); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unrecognized file extension")
	}

	return nil
	return nil
}

func restoreCSV(filename string, storage *postgres, logger logrus.FieldLogger) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return insert(file, filepath.Base(filename), storage, logger)
}

func restoreZIP(filename string, storage *postgres, logger logrus.FieldLogger) error {
	return nil
}

func insert(fileReader io.Reader, filename string, storage *postgres, logger logrus.FieldLogger) error {
	start := time.Now()
	converter := NewConverter(filename)
	if converter == nil {
		logger.Warnf("unrecognized filename: %s", filename)
		return nil
	}

	csvReader := csv.NewReader(fileReader)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	csvReader.TrimLeadingSpace = true

	headers, err := csvReader.Read()
	if err != nil {
		return err
	}

	// TODO: Improve truncate
	if err := truncate(filename, storage, logger); err != nil {
		return err
	}

	line := 2
	limit := 1000
	wg := NewWaitGroup(runtime.NumCPU())
	for {
		rows, err := read(csvReader, headers, limit)
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
				model := converter(*row)
				model.SetID(line)
				models = append(models, model)
				line++

				if (line % limit) == 0 {
					logger.Warnf("Processed %s: %d lines", filename, line)
				}
			}

			if err := storage.db.Insert(&models); err != nil {
				logger.Error(err)
			}
		}()

		wg.Wait()
	}

	logger.Infof("Processed %s: %d lines took %.2fmin", filename, line, time.Since(start).Minutes())
	return nil
}

func read(reader *csv.Reader, headers []string, limit int) ([]*map[string]string, error) {
	models := []*map[string]string{}
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

		models = append(models, &row)
		total++

		if total >= limit {
			break
		}
	}

	return models, nil
}
