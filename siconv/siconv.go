package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Create(model schema.Model) error
	Get(model schema.Model) error
	List(models interface{}, pagination *storage.Pagination) (int, error)
}

type siconv struct {
	storage storage.Service
	logger  logrus.FieldLogger
}

func (s *siconv) Create(model schema.Model) error {
	return s.storage.Insert(model)
}

func (s *siconv) Get(model schema.Model) error {
	return s.storage.Lookup(model)
}

func (s *siconv) List(models interface{}, pagination *storage.Pagination) (int, error) {
	return s.storage.List(models, pagination)
}

func New(storage storage.Service, logger logrus.FieldLogger) Service {
	return &siconv{
		storage: storage,
		logger:  logger.WithField("component", "siconv"),
	}
}
