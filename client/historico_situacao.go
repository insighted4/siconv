package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreateHistoricoSituacao(historicoSituacao *schema.HistoricoSituacao) (string, error) {
	url := path.Join(s.prefix, "historico-situacoes")
	return s.post(historicoSituacao, url)
}

func (s *Client) GetHistoricoSituacao(id string) (*schema.HistoricoSituacao, error) {
	url := path.Join(s.prefix, "historico-situacoes", id)
	var model schema.HistoricoSituacao
	_, err := s.get(&model, url, nil)
	return &model, err
}
