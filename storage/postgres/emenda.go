package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (dao *postgres) CreateEmenda(emenda *schema.Emenda) (string, error) {
	if _, err := dao.db.Model(emenda).Insert(); err != nil {
		return "", err
	}

	return emenda.ID, nil
}

func (dao *postgres) GetEmenda(id string) (*schema.Emenda, error) {
	var model schema.Emenda
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListEmenda(pagination *siconv.Pagination) ([]*schema.Emenda, int, error) {
	models := []*schema.Emenda{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
