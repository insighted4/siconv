package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreateTermoAditivo(termoAditivo *schema.TermoAditivo) (string, error) {
	if _, err := dao.db.Model(termoAditivo).Insert(); err != nil {
		return "", err
	}

	return termoAditivo.ID, nil
}

func (dao *postgres) GetTermoAditivo(id string) (*schema.TermoAditivo, error) {
	var model schema.TermoAditivo
	_, err := dao.get(&model, id)

	return &model, err
}
