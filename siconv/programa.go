package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreatePrograma(programa *schema.Programa) (string, error) {
	return s.dao.CreatePrograma(programa)
}

func (s *service) GetPrograma(id string) (*schema.Programa, error) {
	return s.dao.GetPrograma(id)
}

func (s *service) ListPrograma(idPrograma string, pagination *Pagination) ([]*schema.Programa, int, error) {
	return s.dao.ListPrograma(idPrograma, pagination)
}
