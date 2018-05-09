package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
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

func (s *Client) ListEtapaCronoFisico(pagination *storage.Pagination) ([]*schema.EtapaCronoFisico, int, error) {
	url := path.Join(s.prefix, "etapa-crono-fisicos")
	var models []*schema.EtapaCronoFisico
	total, err := s.get(&models, url, nil)
	return models, total, err
}
