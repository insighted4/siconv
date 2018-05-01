package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateConvenio(convenio *schema.Convenio) (string, error) {
	return s.dao.CreateConvenio(convenio)
}

func (s *service) GetConvenio(id string) (*schema.Convenio, error) {
	return s.dao.GetConvenio(id)
}
