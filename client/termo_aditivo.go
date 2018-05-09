package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
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

func (s *Client) ListTermoAditivo(pagination *storage.Pagination) ([]*schema.TermoAditivo, int, error) {
	url := path.Join(s.prefix, "termo-aditivos")
	var models []*schema.TermoAditivo
	total, err := s.get(&models, url, nil)
	return models, total, err
}
