package postgres

import (
	"github.com/insighted4/siconv/schema"
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
