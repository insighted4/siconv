package postgres

import (
	"github.com/insighted4/siconv/schema"
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
