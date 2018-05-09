package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (dao *postgres) CreateEmpenho(empenho *schema.Empenho) (string, error) {
	if _, err := dao.db.Model(empenho).Insert(); err != nil {
		return "", err
	}

	return empenho.ID, nil
}

func (dao *postgres) GetEmpenho(id string) (*schema.Empenho, error) {
	var model schema.Empenho
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListEmpenho(pagination *storage.Pagination) ([]*schema.Empenho, int, error) {
	models := []*schema.Empenho{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
