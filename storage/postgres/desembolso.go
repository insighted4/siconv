package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreateDesembolso(desembolso *schema.Desembolso) (string, error) {
	if _, err := dao.db.Model(desembolso).Insert(); err != nil {
		return "", err
	}

	return desembolso.ID, nil
}

func (dao *postgres) GetDesembolso(id string) (*schema.Desembolso, error) {
	var model schema.Desembolso
	_, err := dao.get(&model, id)

	return &model, err
}
