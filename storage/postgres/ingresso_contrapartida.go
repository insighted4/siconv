package postgres

import (
	"github.com/insighted4/siconv/schema"
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
