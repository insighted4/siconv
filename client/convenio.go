package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *Client) CreateConvenio(convenio *schema.Convenio) (string, error) {
	url := path.Join(s.prefix, "convenios")
	return s.post(convenio, url)
}

func (s *Client) GetConvenio(id string) (*schema.Convenio, error) {
	url := path.Join(s.prefix, "convenios", id)
	var model schema.Convenio
	_, err := s.get(&model, url, nil)
	return &model, err
}

func (s *Client) ListConvenio(pagination *storage.Pagination) ([]*schema.Convenio, int, error) {
	url := path.Join(s.prefix, "convenios")
	var models []*schema.Convenio
	total, err := s.get(&models, url, nil)
	return models, total, err
}
