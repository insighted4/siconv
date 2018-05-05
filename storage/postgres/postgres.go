package postgres

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
	"github.com/sirupsen/logrus"
)

type postgres struct {
	db     *pg.DB
	logger logrus.FieldLogger
}

func (dao *postgres) Check() error {
	return nil
}

func (dao *postgres) insert(model schema.Model) (string, error) {
	if _, err := dao.db.Model(model).Insert(); err != nil {
		return "", err
	}

	return model.GetID(), nil
}

func (dao *postgres) get(model schema.Model, id string) (schema.Model, error) {
	err := dao.db.Model(model).Where("id = ?", id).Select()
	if err == pg.ErrNoRows {
		return model, siconv.ErrNotFound
	}

	return model, err
}

func (dao *postgres) query(models interface{}, sql string, countSql string, pagination *siconv.Pagination, params ...interface{}) (interface{}, int, error) {
	if pagination == nil {
		pagination = siconv.NewPagination(siconv.Limit, 0)
	}

	params = append(params, pagination.Limit)
	params = append(params, pagination.Offset)

	paginatedSql := sql + " LIMIT ? OFFSET ?"
	if _, err := dao.db.Query(models, paginatedSql, params...); err != nil && err != pg.ErrNoRows {
		return models, 0, err
	}

	var count int
	if _, err := dao.db.Query(&count, countSql, params...); err != nil && err != pg.ErrNoRows {
		return models, 0, err
	}

	return models, count, nil
}

func (dao *postgres) selectAndCount(model interface{}, pagination *siconv.Pagination) (interface{}, int, error) {
	if pagination == nil {
		pagination = siconv.NewPagination(siconv.Limit, 0)
	}

	count, err := dao.db.Model(model).Limit(pagination.Limit).Offset(pagination.Offset).Order("id").SelectAndCount()
	if err != nil && err != pg.ErrNoRows {
		return nil, 0, err
	}

	return model, count, nil
}

func New(options *pg.Options, logger logrus.FieldLogger) siconv.Service {
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
