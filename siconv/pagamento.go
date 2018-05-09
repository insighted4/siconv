package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreatePagamento(pagamento *schema.Pagamento) (string, error) {
	return s.dao.CreatePagamento(pagamento)
}

func (s *service) GetPagamento(id string) (*schema.Pagamento, error) {
	return s.dao.GetPagamento(id)
}

func (s *service) ListPagamento(pagination *storage.Pagination) ([]*schema.Pagamento, int, error) {
	return s.dao.ListPagamento(pagination)
}
