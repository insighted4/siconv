package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (dao *postgres) CreateIngressoContrapartida(ingressoContraPartida *schema.IngressoContrapartida) (string, error) {
	if _, err := dao.db.Model(ingressoContraPartida).Insert(); err != nil {
		return "", err
	}

	return ingressoContraPartida.ID, nil
}

func (dao *postgres) GetIngressoContrapartida(id string) (*schema.IngressoContrapartida, error) {
	var model schema.IngressoContrapartida
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListIngressoContrapartida(pagination *storage.Pagination) ([]*schema.IngressoContrapartida, int, error) {
	models := []*schema.IngressoContrapartida{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
