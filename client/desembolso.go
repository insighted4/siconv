package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *Client) CreateDesembolso(desembolso *schema.Desembolso) (string, error) {
	url := path.Join(s.prefix, "desembolsos")
	return s.post(desembolso, url)
}

func (s *Client) GetDesembolso(id string) (*schema.Desembolso, error) {
	url := path.Join(s.prefix, "desembolsos", id)
	var model schema.Desembolso
	_, err := s.get(&model, url, nil)
	return &model, err
}

func (s *Client) ListDesembolso(pagination *storage.Pagination) ([]*schema.Desembolso, int, error) {
	url := path.Join(s.prefix, "desembolsos")
	var models []*schema.Desembolso
	total, err := s.get(&models, url, nil)
	return models, total, err
}
