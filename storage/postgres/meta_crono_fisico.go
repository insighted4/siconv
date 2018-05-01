package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreateMetaCronoFisico(metaCronoFisico *schema.MetaCronoFisico) (string, error) {
	if _, err := dao.db.Model(metaCronoFisico).Insert(); err != nil {
		return "", err
	}

	return metaCronoFisico.ID, nil
}

func (dao *postgres) GetMetaCronoFisico(id string) (*schema.MetaCronoFisico, error) {
	var model schema.MetaCronoFisico
	_, err := dao.get(&model, id)

	return &model, err
}
