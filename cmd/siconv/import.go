package main

import (
	"archive/zip"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"

	"github.com/insighted4/siconv/client"
	"github.com/insighted4/siconv/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func commandImport() *cobra.Command {
	var (
		file       string
		addressURL string
		token      string
		logLevel   string
		logFormat  string
	)
	cmd := cobra.Command{
		Use:     "import",
		Short:   "Import database from an archive file",
		Example: fmt.Sprintf("%s import --address https://siconv-0a11.lab.insighted4.io --file sample/siconv.zip --token B8v83HQh", ShortDescription),
		Run: func(cmd *cobra.Command, args []string) {
			parsedURL, err := url.Parse(viper.GetString("address"))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			cli := client.New(parsedURL.String(), viper.GetString("token"))

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

			logger, err := server.NewLogger(viper.GetString("log_level"), viper.GetString("log_format"))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			start := time.Now()
			if err := importFile(filename, cli, logger); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			logger.Infof("Total: %.2fmin", time.Since(start).Minutes())
		},
	}

	cmd.Flags().StringVar(&file, "file", "siconv.zip", "ZIP or CSV file")
	viper.BindPFlag("file", cmd.Flags().Lookup("file"))

	cmd.Flags().StringVar(&addressURL, "address", "http://localhost:8080", "Server address")
	viper.BindPFlag("address", cmd.Flags().Lookup("address"))

	cmd.Flags().StringVar(&token, "token", "", "Authentication token")
	viper.BindPFlag("token", cmd.Flags().Lookup("token"))

	cmd.Flags().StringVar(&logLevel, "log-level", "info", "Logger level")
	viper.BindPFlag("log_level", cmd.Flags().Lookup("log-level"))

	cmd.Flags().StringVar(&logFormat, "log-format", "text", "Logger format")
	viper.BindPFlag("log_format", cmd.Flags().Lookup("log-format"))

	return &cmd
}

func importFile(filename string, cli *client.Client, logger logrus.FieldLogger) error {
	switch filepath.Ext(filename) {
	case ".csv":
		file, err := os.Open(filename)
		if err != nil {
			return err
		}

		defer file.Close()

		logger.Infof("Uploading %s", filename)
		return cli.Upload(filepath.Base(filename), file)
	case ".zip":
		logger.Infof("Extracting %s", filename)
		zipReader, err := zip.OpenReader(filename)
		if err != nil {
			return err
		}
		defer zipReader.Close()

		for _, file := range zipReader.File {
			if file.FileInfo().IsDir() {
				continue
			}

			fileReader, err := file.Open()
			if err != nil {
				return err
			}

			filename := file.FileInfo().Name()

			logger.Infof("Uploading %s", filename)
			if err := cli.Upload(filename, fileReader); err != nil {
				fileReader.Close()
				return err
			}

			fileReader.Close()
		}
	default:
		return fmt.Errorf("unrecognized file extension")
	}

	return nil
}
