package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreateEmpenho(empenho *schema.Empenho) (string, error) {
	if _, err := dao.db.Model(empenho).Insert(); err != nil {
		return "", err
	}

	return empenho.ID, nil
}

func (dao *postgres) GetEmpenho(id string) (*schema.Empenho, error) {
	var model schema.Empenho
	_, err := dao.get(&model, id)

	return &model, err
}
