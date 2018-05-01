package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
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