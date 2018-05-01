package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateEmpenhoDesembolso(empenhoDesembolso *schema.EmpenhoDesembolso) (string, error) {
	return s.dao.CreateEmpenhoDesembolso(empenhoDesembolso)
}
