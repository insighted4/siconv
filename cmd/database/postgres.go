package main

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

type postgres struct {
	db     *pg.DB
	logger logrus.FieldLogger
}

func NewPostgres(options *pg.Options, logger logrus.FieldLogger) *postgres {
	logger = logger.WithField("component", "postgres")

	db := pg.Connect(options)
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}

		logger.WithField("query", query).
			WithField("latency", time.Since(event.StartTime)).
			Debugf("Query processed")
	})

	return &postgres{
		db:     db,
		logger: logger,
	}
}
