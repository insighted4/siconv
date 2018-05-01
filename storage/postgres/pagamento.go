package postgres

import (
	"github.com/insighted4/siconv/schema"
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
