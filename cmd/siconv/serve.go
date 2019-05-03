package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/go-pg/pg"
	"github.com/insighted4/siconv/server"
	"github.com/insighted4/siconv/storage/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func commandServe() *cobra.Command {
	var (
		databaseURL string
		token       string
		logLevel    string
		logFormat   string
	)
	cmd := cobra.Command{
		Use:     "serve",
		Short:   "Start HTTP server",
		Example: fmt.Sprintf("%s serve --database postgres://user:secret@localhost:5432/siconv --token B8v83HQh", ShortDescription),
		Run: func(cmd *cobra.Command, args []string) {
			u, err := url.Parse(viper.GetString("database_url"))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			pwd, _ := u.User.Password()
			pgOptions := &pg.Options{
				Addr:     u.Host,
				Database: strings.Replace(u.RequestURI(), "/", "", 1),
				User:     u.User.Username(),
				Password: pwd,
			}

			if err := serve(viper.GetString("token"), pgOptions, viper.GetString("log_level"), viper.GetString("log_format")); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVar(&databaseURL, "database-url", "postgres://localhost:5432/siconv", "Database connection string")
	viper.BindPFlag("database_url", cmd.Flags().Lookup("database-url"))

	cmd.Flags().StringVar(&token, "token", "", "Authentication token")
	viper.BindPFlag("token", cmd.Flags().Lookup("token"))

	cmd.Flags().StringVar(&logLevel, "log-level", "info", "Logger level")
	viper.BindPFlag("log_level", cmd.Flags().Lookup("log-level"))

	cmd.Flags().StringVar(&logFormat, "log-format", "text", "Logger format")
	viper.BindPFlag("log_format", cmd.Flags().Lookup("log-format"))

	return &cmd
}

func serve(token string, pgOptions *pg.Options, logLevel string, logFormat string) error {
	logger, err := server.NewLogger(logLevel, logFormat)
	if err != nil {
		return err
	}

	pg := postgres.New(pgOptions, logger)
	cfg := server.Config{
		Token:   token,
		Storage: pg,
		Logger:  logger,
	}

	logger.Info("Starting SICONV API")
	logger.Infof("Authorization Token: %s", token)
	logger.Infof("Database: postgres://%s/%s", pgOptions.Addr, pgOptions.Database)
	logger.Infof("Logger Level: %s", logLevel)
	logger.Infof("Logger Format: %s", logFormat)
	srv, err := server.New(cfg)
	if err != nil {
		return err
	}

	return srv.RunHTTPServer("0.0.0.0:8080")
}
