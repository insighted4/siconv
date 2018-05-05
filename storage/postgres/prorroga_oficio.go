package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (dao *postgres) CreateProrrogaOficio(prorrogaOficio *schema.ProrrogaOficio) (string, error) {
	if _, err := dao.db.Model(prorrogaOficio).Insert(); err != nil {
		return "", err
	}

	return prorrogaOficio.ID, nil
}

func (dao *postgres) GetProrrogaOficio(id string) (*schema.ProrrogaOficio, error) {
	var model schema.ProrrogaOficio
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListProrrogaOficio(pagination *siconv.Pagination) ([]*schema.ProrrogaOficio, int, error) {
	models := []*schema.ProrrogaOficio{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
