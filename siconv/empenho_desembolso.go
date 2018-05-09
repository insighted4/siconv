package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateEmpenhoDesembolso(empenhoDesembolso *schema.EmpenhoDesembolso) (string, error) {
	return s.dao.CreateEmpenhoDesembolso(empenhoDesembolso)
}

func (s *service) ListEmpenhoDesembolso(pagination *storage.Pagination) ([]*schema.EmpenhoDesembolso, int, error) {
	return s.dao.ListEmpenhoDesembolso(pagination)
}
