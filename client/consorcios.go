package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreateConsorcio(consorcio *schema.Consorcio) (string, error) {
	url := path.Join(s.prefix, "consorcios")
	return s.post(consorcio, url)
}

func (s *Client) GetConsorcio(id string) (*schema.Consorcio, error) {
	url := path.Join(s.prefix, "consorcios", id)
	var model schema.Consorcio
	_, err := s.get(&model, url, nil)
	return &model, err
}