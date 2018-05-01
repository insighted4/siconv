package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreateEmpenhoDesembolso(empenhoDesembolso *schema.EmpenhoDesembolso) (string, error) {
	if _, err := dao.db.Model(empenhoDesembolso).Insert(); err != nil {
		return "", err
	}

	return empenhoDesembolso.ID, nil
}
