package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *Client) CreateMetaCronoFisico(metaCronoFisico *schema.MetaCronoFisico) (string, error) {
	url := path.Join(s.prefix, "meta-crono-fisicos")
	return s.post(metaCronoFisico, url)
}

func (s *Client) GetMetaCronoFisico(id string) (*schema.MetaCronoFisico, error) {
	url := path.Join(s.prefix, "meta-crono-fisicos", id)
	var model schema.MetaCronoFisico
	_, err := s.get(&model, url, nil)
	return &model, err
}

func (s *Client) ListMetaCronoFisico(pagination *storage.Pagination) ([]*schema.MetaCronoFisico, int, error) {
	url := path.Join(s.prefix, "meta-crono-fisicos")
	var models []*schema.MetaCronoFisico
	total, err := s.get(&models, url, nil)
	return models, total, err
}
