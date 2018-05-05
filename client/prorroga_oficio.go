package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *Client) CreateProrrogaOficio(prorrogaOficio *schema.ProrrogaOficio) (string, error) {
	url := path.Join(s.prefix, "prorroga-oficios")
	return s.post(prorrogaOficio, url)
}

func (s *Client) GetProrrogaOficio(id string) (*schema.ProrrogaOficio, error) {
	url := path.Join(s.prefix, "prorroga-oficios", id)
	var model schema.ProrrogaOficio
	_, err := s.get(&model, url, nil)
	return &model, err
}

func (s *Client) ListProrrogaOficio(pagination *siconv.Pagination) ([]*schema.ProrrogaOficio, int, error) {
	url := path.Join(s.prefix, "prorroga-oficios")
	var models []*schema.ProrrogaOficio
	total, err := s.get(&models, url, nil)
	return models, total, err
}
