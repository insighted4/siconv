package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreatePagamento(pagamento *schema.Pagamento) (string, error) {
	url := path.Join(s.prefix, "pagamentos")
	return s.post(pagamento, url)
}

func (s *Client) GetPagamento(id string) (*schema.Pagamento, error) {
	url := path.Join(s.prefix, "pagamentos", id)
	var model schema.Pagamento
	_, err := s.get(&model, url, nil)
	return &model, err
}
