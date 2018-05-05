package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (dao *postgres) CreateOBTVConvenente(obtvConvenente *schema.OBTVConvenente) (string, error) {
	if _, err := dao.db.Model(obtvConvenente).Insert(); err != nil {
		return "", err
	}

	return obtvConvenente.ID, nil
}

func (dao *postgres) GetOBTVConvenente(id string) (*schema.OBTVConvenente, error) {
	var model schema.OBTVConvenente
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListOBTVConvenente(pagination *siconv.Pagination) ([]*schema.OBTVConvenente, int, error) {
	models := []*schema.OBTVConvenente{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
