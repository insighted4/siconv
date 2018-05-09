package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateEmpenho(empenho *schema.Empenho) (string, error) {
	return s.dao.CreateEmpenho(empenho)
}

func (s *service) GetEmpenho(id string) (*schema.Empenho, error) {
	return s.dao.GetEmpenho(id)
}

func (s *service) ListEmpenho(pagination *storage.Pagination) ([]*schema.Empenho, int, error) {
	return s.dao.ListEmpenho(pagination)
}
