package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateOBTVConvenente(obtvConvenente *schema.OBTVConvenente) (string, error) {
	return s.dao.CreateOBTVConvenente(obtvConvenente)
}

func (s *service) GetOBTVConvenente(id string) (*schema.OBTVConvenente, error) {
	return s.dao.GetOBTVConvenente(id)
}

func (s *service) ListOBTVConvenente(pagination *storage.Pagination) ([]*schema.OBTVConvenente, int, error) {
	return s.dao.ListOBTVConvenente(pagination)
}
