package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/insighted4/siconv/client"
	"github.com/insighted4/siconv/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func commandSync() *cobra.Command {
	var (
		file         string
		serverURL    string
		token        string
		loggerLevel  string
		loggerFormat string
		maxWorkers   int
		maxQueueSize int
	)

	cmd := cobra.Command{
		Use:     "sync",
		Short:   "Sync database with downloaded zip file",
		Example: "updater sync",
		Run: func(cmd *cobra.Command, args []string) {
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

			logger, err := server.NewLogger(loggerLevel, loggerFormat)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			cli, err := client.New(serverURL, token)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}

			start := time.Now()
			logger = logger.WithField("component", "updater")
			if err := syncFile(filename, cli, logger, maxWorkers, maxQueueSize); err != nil {
				logger.Error(err)
				os.Exit(1)
			}
			logger.Infof("Total: %.2fmin", time.Since(start).Minutes())
		},
	}

	cmd.Flags().StringVar(&file, "file", "siconv.zip", "ZIP or CSV file")
	cmd.Flags().StringVar(&serverURL, "server", "http://localhost", "API Server URL")
	cmd.Flags().StringVar(&token, "token", "", "Authentication token")
	cmd.Flags().IntVar(&maxWorkers, "max-workers", 2, "Max workers")
	cmd.Flags().IntVar(&maxQueueSize, "max-queue-size", 4, "Max queue size")
	cmd.Flags().StringVar(&loggerLevel, "log-level", "info", "Logger level")
	cmd.Flags().StringVar(&loggerFormat, "log-format", "text", "Logger format")

	return &cmd
}

func syncFile(filename string, cli *client.Client, logger logrus.FieldLogger, maxWorkers int, maxQueueSize int) error {
	// Initialize Workers
	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers, logger)
	dispatcher.Run()

	logger.Infof("Updating database %s", filename)

	switch filepath.Ext(filename) {
	case ".zip":
		if err := syncZIP(filename, cli, logger, jobQueue); err != nil {
			return err
		}
	case ".csv":
		if err := syncCSV(filename, cli, logger, jobQueue); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unrecognized file extension")
	}

	return nil
}

func syncZIP(filename string, cli *client.Client, logger logrus.FieldLogger, jobQueue chan Job) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}

		logger.Infof("Processing %s", file.FileInfo().Name())
		if err := processZIP(file, cli, logger, jobQueue); err != nil {
			return err
		}
	}

	return nil
}

func syncCSV(filename string, cli *client.Client, logger logrus.FieldLogger, jobQueue chan Job) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	filename = filepath.Base(filename)
	logger.Infof("Processing %s", filename)
	if err := processCSV(file, filename, cli, logger, jobQueue); err != nil {
		return err
	}

	return nil
}

func processZIP(file *zip.File, cli *client.Client, logger logrus.FieldLogger, jobQueue chan Job) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	filename := file.FileInfo().Name()
	return processCSV(fileReader, filename, cli, logger, jobQueue)
}

type HandlerFunc func(r map[string]string) (string, error)

func processCSV(fileReader io.Reader, filename string, cli *client.Client, logger logrus.FieldLogger, jobQueue chan Job) error {
	start := time.Now()
	handler := NewHandler(filename, cli)
	if handler == nil {
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

	line := 2
	var wg sync.WaitGroup
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		row := map[string]string{
			"file_ref": filename,
			"line_ref": fmt.Sprintf("%d", line),
		}
		for index, header := range headers {
			h := strings.TrimSpace(strings.ToLower(header))
			h = strings.Trim(h, "\xef\xbb\xbf") // Remove infamous BOM

			v := strings.TrimSpace(record[index])
			row[h] = v
		}

		wg.Add(1)
		jobQueue <- NewUpdate(filename, line, row, handler, logger, &wg)

		if (line % 1000) == 0 {
			wg.Wait()
			logger.Infof("Processed %s: %d lines", filename, line)
		}

		line++
	}

	wg.Wait()
	logger.Infof("Processed %s: %d lines took %.2fmin", filename, line, time.Since(start).Minutes())
	return nil
}

func NewHandler(filename string, c *client.Client) HandlerFunc {
	var handler HandlerFunc
	switch filename {
	case "siconv_consorcios.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateConsorcio(NewConsorcio(row))
		}
	case "siconv_convenio.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateConvenio(NewConvenio(row))
		}
	case "siconv_desembolso.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateDesembolso(NewDesembolso(row))
		}
	case "siconv_emenda.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateEmenda(NewEmenda(row))
		}
	case "siconv_empenho.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateEmpenho(NewEmpenho(row))
		}
	case "siconv_empenho_desembolso.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateEmpenhoDesembolso(NewEmpenhoDesembolso(row))
		}
	case "siconv_etapa_crono_fisico.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateEtapaCronoFisico(NewEtapaCronoFisico(row))
		}
	case "siconv_historico_situacao.csv":
		handler = func(row map[string]string) (string, error) {
			model := NewHistoricoSituacao(row)
			if model.DIA_HISTORICO_SIT == nil || model.DIAS_HISTORICO_SIT == 0 {
				return "(ignored)", nil
			}
			return c.CreateHistoricoSituacao(model)
		}
	case "siconv_ingresso_contrapartida.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateIngressoContrapartida(NewIngressoContrapartida(row))
		}
	case "siconv_meta_crono_fisico.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateMetaCronoFisico(NewMetaCronoFisico(row))
		}
	case "siconv_obtv_convenente.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateOBTVConvenente(NewOBTVConvenente(row))
		}
	case "siconv_pagamento.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreatePagamento(NewPagamento(row))
		}
	case "siconv_plano_aplicacao_detalhado.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreatePlanoAplicacaoDetalhado(NewPlanoAplicacaoDetalhado(row))
		}
	case "siconv_programa.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreatePrograma(NewPrograma(row))
		}
	case "siconv_programa_proposta.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateProgramaProposta(NewProgramaProposta(row))
		}
	case "siconv_proponentes.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateProponente(NewProponente(row))
		}
	case "siconv_proposta.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateProposta(NewProposta(row))
		}
	case "siconv_prorroga_oficio.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateProrrogaOficio(NewProrrogaOficio(row))
		}
	case "siconv_termo_aditivo.csv":
		handler = func(row map[string]string) (string, error) {
			return c.CreateTermoAditivo(NewTermoAditivo(row))
		}
	}

	return handler
}
