package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateMetaCronoFisico(metaCronoFisico *schema.MetaCronoFisico) (string, error) {
	return s.dao.CreateMetaCronoFisico(metaCronoFisico)
}

func (s *service) GetMetaCronoFisico(id string) (*schema.MetaCronoFisico, error) {
	return s.dao.GetMetaCronoFisico(id)
}

func (s *service) ListMetaCronoFisico(pagination *Pagination) ([]*schema.MetaCronoFisico, int, error) {
	return s.dao.ListMetaCronoFisico(pagination)
}
