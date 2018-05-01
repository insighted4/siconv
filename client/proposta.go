package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreateProposta(proposta *schema.Proposta) (string, error) {
	url := path.Join(s.prefix, "propostas")
	return s.post(proposta, url)
}

func (s *Client) GetProposta(id string) (*schema.Proposta, error) {
	url := path.Join(s.prefix, "propostas", id)
	var model schema.Proposta
	_, err := s.get(&model, url, nil)
	return &model, err
}
