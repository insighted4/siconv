package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateConvenio(convenio *schema.Convenio) (string, error) {
	return s.dao.CreateConvenio(convenio)
}

func (s *service) GetConvenio(id string) (*schema.Convenio, error) {
	return s.dao.GetConvenio(id)
}

func (s *service) ListConvenio(pagination *storage.Pagination) ([]*schema.Convenio, int, error) {
	return s.dao.ListConvenio(pagination)
}
