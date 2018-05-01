package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado *schema.PlanoAplicacaoDetalhado) (string, error) {
	return s.dao.CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado)
}

func (s *service) GetPlanoAplicacaoDetalhado(id string) (*schema.PlanoAplicacaoDetalhado, error) {
	return s.dao.GetPlanoAplicacaoDetalhado(id)
}