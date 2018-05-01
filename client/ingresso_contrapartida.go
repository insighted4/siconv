package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreateIngressoContrapartida(ingressoContraPartida *schema.IngressoContrapartida) (string, error) {
	url := path.Join(s.prefix, "ingresso-contrapartidas")
	return s.post(ingressoContraPartida, url)
}

func (s *Client) GetIngressoContrapartida(id string) (*schema.IngressoContrapartida, error) {
	url := path.Join(s.prefix, "ingresso-contrapartidas", id)
	var model schema.IngressoContrapartida
	_, err := s.get(&model, url, nil)
	return &model, err
}