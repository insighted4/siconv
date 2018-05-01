package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado *schema.PlanoAplicacaoDetalhado) (string, error) {
	url := path.Join(s.prefix, "plano-aplicacao-detalhados")
	return s.post(planoAplicacaoDetalhado, url)
}

func (s *Client) GetPlanoAplicacaoDetalhado(id string) (*schema.PlanoAplicacaoDetalhado, error) {
	url := path.Join(s.prefix, "plano-aplicacao-detalhados", id)
	var model schema.PlanoAplicacaoDetalhado
	_, err := s.get(&model, url, nil)
	return &model, err
}