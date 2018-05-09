package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *Client) CreateOBTVConvenente(obtvConvenente *schema.OBTVConvenente) (string, error) {
	url := path.Join(s.prefix, "obtv-convenentes")
	return s.post(obtvConvenente, url)
}

func (s *Client) GetOBTVConvenente(id string) (*schema.OBTVConvenente, error) {
	url := path.Join(s.prefix, "obtv-convenentes", id)
	var model schema.OBTVConvenente
	_, err := s.get(&model, url, nil)
	return &model, err
}

func (s *Client) ListOBTVConvenente(pagination *storage.Pagination) ([]*schema.OBTVConvenente, int, error) {
	url := path.Join(s.prefix, "obtv-convenentes")
	var models []*schema.OBTVConvenente
	total, err := s.get(&models, url, nil)
	return models, total, err
}
