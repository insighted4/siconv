package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateProrrogaOficio(prorrogaOficio *schema.ProrrogaOficio) (string, error) {
	return s.dao.CreateProrrogaOficio(prorrogaOficio)
}

func (s *service) GetProrrogaOficio(id string) (*schema.ProrrogaOficio, error) {
	return s.dao.GetProrrogaOficio(id)
}

func (s *service) ListProrrogaOficio(pagination *storage.Pagination) ([]*schema.ProrrogaOficio, int, error) {
	return s.dao.ListProrrogaOficio(pagination)
}
