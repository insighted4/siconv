package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateConsorcio(consorcio *schema.Consorcio) (string, error) {
	return s.dao.CreateConsorcio(consorcio)
}

func (s *service) GetConsorcio(id string) (*schema.Consorcio, error) {
	return s.dao.GetConsorcio(id)
}

func (s *service) ListConsorcio(pagination *Pagination) ([]*schema.Consorcio, int, error) {
	return s.dao.ListConsorcio(pagination)
}
