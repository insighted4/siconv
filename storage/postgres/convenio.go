package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (dao *postgres) CreateConvenio(convenio *schema.Convenio) (string, error) {
	if _, err := dao.db.Model(convenio).Insert(); err != nil {
		return "", err
	}

	return convenio.ID, nil
}

func (dao *postgres) GetConvenio(id string) (*schema.Convenio, error) {
	var model schema.Convenio
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListConvenio(pagination *siconv.Pagination) ([]*schema.Convenio, int, error) {
	models := []*schema.Convenio{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
