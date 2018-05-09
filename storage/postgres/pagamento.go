package postgres

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (dao *postgres) CreatePagamento(pagamento *schema.Pagamento) (string, error) {
	if _, err := dao.db.Model(pagamento).Insert(); err != nil {
		return "", err
	}

	return pagamento.ID, nil
}

func (dao *postgres) GetPagamento(id string) (*schema.Pagamento, error) {
	var model schema.Pagamento
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListPagamento(pagination *storage.Pagination) ([]*schema.Pagamento, int, error) {
	models := []*schema.Pagamento{nil}
	_, count, err := dao.selectAndCount(&models, pagination)
	return models, count, err
}
