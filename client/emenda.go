package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreateEmenda(emenda *schema.Emenda) (string, error) {
	url := path.Join(s.prefix, "emendas")
	return s.post(emenda, url)
}

func (s *Client) GetEmenda(id string) (*schema.Emenda, error) {
	url := path.Join(s.prefix, "emendas", id)
	var model schema.Emenda
	_, err := s.get(&model, url, nil)
	return &model, err
}
