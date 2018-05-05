package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *Client) CreateEmpenho(empenho *schema.Empenho) (string, error) {
	url := path.Join(s.prefix, "empenhos")
	return s.post(empenho, url)
}

func (s *Client) GetEmpenho(id string) (*schema.Empenho, error) {
	url := path.Join(s.prefix, "empenhos", id)
	var model schema.Empenho
	_, err := s.get(&model, url, nil)
	return &model, err
}

func (s *Client) ListEmpenho(pagination *siconv.Pagination) ([]*schema.Empenho, int, error) {
	url := path.Join(s.prefix, "empenhos")
	var models []*schema.Empenho
	total, err := s.get(&models, url, nil)
	return models, total, err
}
