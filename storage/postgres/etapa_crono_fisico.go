package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (dao *postgres) CreateEtapaCronoFisico(etapaCronoFisico *schema.EtapaCronoFisico) (string, error) {
	if _, err := dao.db.Model(etapaCronoFisico).Insert(); err != nil {
		return "", err
	}

	return etapaCronoFisico.ID, nil
}

func (dao *postgres) GetEtapaCronoFisico(id string) (*schema.EtapaCronoFisico, error) {
	var model schema.EtapaCronoFisico
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListEtapaCronoFisico(pagination *storage.Pagination) ([]*schema.EtapaCronoFisico, int, error) {
	models := []*schema.EtapaCronoFisico{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
