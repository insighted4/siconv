package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *Client) CreateProponente(proponente *schema.Proponente) (string, error) {
	url := path.Join(s.prefix, "proponentes")
	return s.post(proponente, url)
}

func (s *Client) GetProponente(id string) (*schema.Proponente, error) {
	url := path.Join(s.prefix, "proponentes", id)
	var model schema.Proponente
	_, err := s.get(&model, url, nil)
	return &model, err
}

func (s *Client) ListProponente(pagination *siconv.Pagination) ([]*schema.Proponente, int, error) {
	url := path.Join(s.prefix, "proponentes")
	var models []*schema.Proponente
	total, err := s.get(&models, url, nil)
	return models, total, err
}
