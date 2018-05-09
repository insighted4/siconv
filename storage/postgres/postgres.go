package postgres

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
	"github.com/sirupsen/logrus"
)

type postgres struct {
	db     *pg.DB
	logger logrus.FieldLogger
}

func (p *postgres) Check() error {
	// TODO
	return nil
}

func (p *postgres) Insert(model schema.Model) error {
	if model.GetID() == 0 {
		return storage.ErrInvalidID
	}

	if _, err := p.db.Model(model).Insert(); err != nil {
		return err
	}

	return nil
}

func (p *postgres) Lookup(model schema.Model) error {
	err := p.db.Model(model).Where("id = ?", model.GetID()).Select()
	if err == pg.ErrNoRows {
		return storage.ErrNotFound
	}

	return err
}

func (p *postgres) List(models interface{}, pagination *storage.Pagination) (int, error) {
	if pagination == nil {
		pagination = storage.NewPagination(storage.Limit, 0)
	}

	count, err := p.db.Model(models).Limit(pagination.Limit).Order("id").Offset(pagination.Offset).SelectAndCount()
	if err != nil && err != pg.ErrNoRows {
		return 0, err
	}

	return count, nil
}

func (p *postgres) query(models interface{}, sql string, countSql string, pagination *storage.Pagination, params ...interface{}) (interface{}, int, error) {
	if pagination == nil {
		pagination = storage.NewPagination(storage.Limit, 0)
	}

	params = append(params, pagination.Limit)
	params = append(params, pagination.Offset)

	paginatedSql := sql + " LIMIT ? OFFSET ?"
	if _, err := p.db.Query(models, paginatedSql, params...); err != nil && err != pg.ErrNoRows {
		return models, 0, err
	}

	var count int
	if _, err := p.db.Query(&count, countSql, params...); err != nil && err != pg.ErrNoRows {
		return models, 0, err
	}

	return models, count, nil
}

func New(options *pg.Options, logger logrus.FieldLogger) storage.Service {
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
