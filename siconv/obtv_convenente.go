package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateOBTVConvenente(obtvConvenente *schema.OBTVConvenente) (string, error) {
	return s.dao.CreateOBTVConvenente(obtvConvenente)
}

func (s *service) GetOBTVConvenente(id string) (*schema.OBTVConvenente, error) {
	return s.dao.GetOBTVConvenente(id)
}