package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateEmpenho(empenho *schema.Empenho) (string, error) {
	return s.dao.CreateEmpenho(empenho)
}

func (s *service) GetEmpenho(id string) (*schema.Empenho, error) {
	return s.dao.GetEmpenho(id)
}
