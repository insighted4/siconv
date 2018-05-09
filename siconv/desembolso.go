package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateDesembolso(desembolso *schema.Desembolso) (string, error) {
	return s.dao.CreateDesembolso(desembolso)
}

func (s *service) GetDesembolso(id string) (*schema.Desembolso, error) {
	return s.dao.GetDesembolso(id)
}

func (s *service) ListDesembolso(pagination *storage.Pagination) ([]*schema.Desembolso, int, error) {
	return s.dao.ListDesembolso(pagination)
}
