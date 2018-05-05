package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateTermoAditivo(termoAditivo *schema.TermoAditivo) (string, error) {
	return s.dao.CreateTermoAditivo(termoAditivo)
}

func (s *service) GetTermoAditivo(id string) (*schema.TermoAditivo, error) {
	return s.dao.GetTermoAditivo(id)
}

func (s *service) ListTermoAditivo(pagination *Pagination) ([]*schema.TermoAditivo, int, error) {
	return s.dao.ListTermoAditivo(pagination)
}
