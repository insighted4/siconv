package main

import (
	"fmt"

	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-pg/pg"
	"github.com/insighted4/siconv/server"
	"github.com/insighted4/siconv/storage/postgres"
	"github.com/spf13/cobra"
)

func commandTrucate() *cobra.Command {
	var (
		table       string
		databaseURL string
		logLevel    string
		logFormat   string
	)
	cmd := cobra.Command{
		Use:     "truncate",
		Short:   "Truncate table",
		Example: "siconv truncate",
		Run: func(cmd *cobra.Command, args []string) {
			parsedURL, err := url.Parse(databaseURL)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
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

			storage := postgres.New(pgOptions, logger)
			start := time.Now()
			logger.Warnf("Truncating table %s", table)
			if err := storage.Truncate(table); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			logger.Infof("Total: %.2fmin", time.Since(start).Minutes())
		},
	}

	cmd.Flags().StringVar(&table, "table", "", "Table name")
	cmd.Flags().StringVar(&databaseURL, "database-url", "postgres://localhost:5432/siconv", "Database connection string. (ex.: postgres://user:pass@localhost:5432/siconv")
	cmd.Flags().StringVar(&logLevel, "log-level", "info", "Logger level")
	cmd.Flags().StringVar(&logFormat, "log-format", "text", "Logger format")

	return &cmd
}
