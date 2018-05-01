package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreateHistoricoSituacao(historicoSituacao *schema.HistoricoSituacao) (string, error) {
	if _, err := dao.db.Model(historicoSituacao).Insert(); err != nil {
		return "", err
	}

	return historicoSituacao.ID, nil
}

func (dao *postgres) GetHistoricoSituacao(id string) (*schema.HistoricoSituacao, error) {
	var model schema.HistoricoSituacao
	_, err := dao.get(&model, id)

	return &model, err
}