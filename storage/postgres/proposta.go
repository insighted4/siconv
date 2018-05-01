package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreateProposta(proposta *schema.Proposta) (string, error) {
	if _, err := dao.db.Model(proposta).Insert(); err != nil {
		return "", err
	}

	return proposta.ID, nil
}

func (dao *postgres) GetProposta(id string) (*schema.Proposta, error) {
	var model schema.Proposta
	_, err := dao.get(&model, id)

	return &model, err
}
