package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
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

func (dao *postgres) ListMetaCronoFisico(pagination *storage.Pagination) ([]*schema.MetaCronoFisico, int, error) {
	models := []*schema.MetaCronoFisico{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
