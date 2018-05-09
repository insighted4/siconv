package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (dao *postgres) CreateEmpenhoDesembolso(empenhoDesembolso *schema.EmpenhoDesembolso) (string, error) {
	if _, err := dao.db.Model(empenhoDesembolso).Insert(); err != nil {
		return "", err
	}

	return empenhoDesembolso.ID, nil
}

func (dao *postgres) ListEmpenhoDesembolso(pagination *storage.Pagination) ([]*schema.EmpenhoDesembolso, int, error) {
	models := []*schema.EmpenhoDesembolso{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
