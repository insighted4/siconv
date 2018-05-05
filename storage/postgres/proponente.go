package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (dao *postgres) CreateProponente(proponente *schema.Proponente) (string, error) {
	if _, err := dao.db.Model(proponente).Insert(); err != nil {
		return "", err
	}

	return proponente.ID, nil
}

func (dao *postgres) GetProponente(id string) (*schema.Proponente, error) {
	var model schema.Proponente
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListProponente(pagination *siconv.Pagination) ([]*schema.Proponente, int, error) {
	models := []*schema.Proponente{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
