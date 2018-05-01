package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateEtapaCronoFisico(etapaCronoFisico *schema.EtapaCronoFisico) (string, error) {
	return s.dao.CreateEtapaCronoFisico(etapaCronoFisico)
}

func (s *service) GetEtapaCronoFisico(id string) (*schema.EtapaCronoFisico, error) {
	return s.dao.GetEtapaCronoFisico(id)
}
