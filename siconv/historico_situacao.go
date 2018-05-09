package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateHistoricoSituacao(historicoSituacao *schema.HistoricoSituacao) (string, error) {
	return s.dao.CreateHistoricoSituacao(historicoSituacao)
}

func (s *service) GetHistoricoSituacao(id string) (*schema.HistoricoSituacao, error) {
	return s.dao.GetHistoricoSituacao(id)
}

func (s *service) ListHistoricoSituacao(pagination *storage.Pagination) ([]*schema.HistoricoSituacao, int, error) {
	return s.dao.ListHistoricoSituacao(pagination)
}
