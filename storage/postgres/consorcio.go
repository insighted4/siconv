package postgres

import (
	"github.com/insighted4/siconv/schema"
)

func (dao *postgres) CreateConsorcio(consorcio *schema.Consorcio) (string, error) {
	if _, err := dao.db.Model(consorcio).Insert(); err != nil {
		return "", err
	}

	return consorcio.ID, nil
}

func (dao *postgres) GetConsorcio(id string) (*schema.Consorcio, error) {
	var model schema.Consorcio
	_, err := dao.get(&model, id)

	return &model, err
}
