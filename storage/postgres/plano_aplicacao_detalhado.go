package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado *schema.PlanoAplicacaoDetalhado) (string, error) {
	if _, err := dao.db.Model(planoAplicacaoDetalhado).Insert(); err != nil {
		return "", err
	}

	return planoAplicacaoDetalhado.ID, nil
}

func (dao *postgres) GetPlanoAplicacaoDetalhado(id string) (*schema.PlanoAplicacaoDetalhado, error) {
	var model schema.PlanoAplicacaoDetalhado
	_, err := dao.get(&model, id)

	return &model, err
}
