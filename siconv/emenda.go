package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateEmenda(emenda *schema.Emenda) (string, error) {
	return s.dao.CreateEmenda(emenda)
}

func (s *service) GetEmenda(id string) (*schema.Emenda, error) {
	return s.dao.GetEmenda(id)
}