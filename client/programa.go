package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *Client) CreatePrograma(programa *schema.Programa) (string, error) {
	url := path.Join(s.prefix, "programas")
	return s.post(programa, url)
}

func (s *Client) GetPrograma(id string) (*schema.Programa, error) {
	url := path.Join(s.prefix, "programas", id)
	var model schema.Programa
	_, err := s.get(&model, url, nil)
	return &model, err
}

func (s *Client) ListPrograma(idPrograma string, pagination *siconv.Pagination) ([]*schema.Programa, int, error) {
	url := path.Join(s.prefix, "programas")
	params := map[string]string{"id_programa": idPrograma}
	var models []*schema.Programa
	total, err := s.get(&models, url, params)
	return models, total, err
}
