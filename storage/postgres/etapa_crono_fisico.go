package postgres

import (
	"github.com/insighted4/siconv/schema"
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
