package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
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

func (dao *postgres) ListProposta(pagination *siconv.Pagination) ([]*schema.Proposta, int, error) {
	models := []*schema.Proposta{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
