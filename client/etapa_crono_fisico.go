package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreateEtapaCronoFisico(etapaCronoFisico *schema.EtapaCronoFisico) (string, error) {
	url := path.Join(s.prefix, "etapa-crono-fisicos")
	return s.post(etapaCronoFisico, url)
}

func (s *Client) GetEtapaCronoFisico(id string) (*schema.EtapaCronoFisico, error) {
	url := path.Join(s.prefix, "etapa-crono-fisicos", id)
	var model schema.EtapaCronoFisico
	_, err := s.get(&model, url, nil)
	return &model, err
}
