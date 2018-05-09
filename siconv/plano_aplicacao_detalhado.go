package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado *schema.PlanoAplicacaoDetalhado) (string, error) {
	return s.dao.CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado)
}

func (s *service) GetPlanoAplicacaoDetalhado(id string) (*schema.PlanoAplicacaoDetalhado, error) {
	return s.dao.GetPlanoAplicacaoDetalhado(id)
}

func (s *service) ListPlanoAplicacaoDetalhado(pagination *storage.Pagination) ([]*schema.PlanoAplicacaoDetalhado, int, error) {
	return s.dao.ListPlanoAplicacaoDetalhado(pagination)
}
