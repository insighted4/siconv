package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreateTermoAditivo(termoAditivo *schema.TermoAditivo) (string, error) {
	url := path.Join(s.prefix, "termo-aditivos")
	return s.post(termoAditivo, url)
}

func (s *Client) GetTermoAditivo(id string) (*schema.TermoAditivo, error) {
	url := path.Join(s.prefix, "termo-aditivos", id)
	var model schema.TermoAditivo
	_, err := s.get(&model, url, nil)
	return &model, err
}
