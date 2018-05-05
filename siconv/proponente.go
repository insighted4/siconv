package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateProponente(proponente *schema.Proponente) (string, error) {
	return s.dao.CreateProponente(proponente)
}
func (s *service) GetProponente(id string) (*schema.Proponente, error) {
	return s.dao.GetProponente(id)
}

func (s *service) ListProponente(pagination *Pagination) ([]*schema.Proponente, int, error) {
	return s.dao.ListProponente(pagination)
}
