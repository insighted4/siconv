package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateProposta(proposta *schema.Proposta) (string, error) {
	return s.dao.CreateProposta(proposta)
}

func (s *service) GetProposta(id string) (*schema.Proposta, error) {
	return s.dao.GetProposta(id)
}
