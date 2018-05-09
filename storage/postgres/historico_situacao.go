package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
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

func (dao *postgres) ListHistoricoSituacao(pagination *storage.Pagination) ([]*schema.HistoricoSituacao, int, error) {
	models := []*schema.HistoricoSituacao{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
