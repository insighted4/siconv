package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateMetaCronoFisico(metaCronoFisico *schema.MetaCronoFisico) (string, error) {
	return s.dao.CreateMetaCronoFisico(metaCronoFisico)
}

func (s *service) GetMetaCronoFisico(id string) (*schema.MetaCronoFisico, error) {
	return s.dao.GetMetaCronoFisico(id)
}

func (s *service) ListMetaCronoFisico(pagination *storage.Pagination) ([]*schema.MetaCronoFisico, int, error) {
	return s.dao.ListMetaCronoFisico(pagination)
}
