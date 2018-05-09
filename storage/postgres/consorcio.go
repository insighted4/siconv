package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
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

func (dao *postgres) ListConsorcio(pagination *storage.Pagination) ([]*schema.Consorcio, int, error) {
	models := []*schema.Consorcio{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
