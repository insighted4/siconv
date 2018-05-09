package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *Client) CreateEmpenhoDesembolso(empenhoDesembolso *schema.EmpenhoDesembolso) (string, error) {
	url := path.Join(s.prefix, "empenho-desembolsos")
	return s.post(empenhoDesembolso, url)
}

func (s *Client) ListEmpenhoDesembolso(pagination *storage.Pagination) ([]*schema.EmpenhoDesembolso, int, error) {
	url := path.Join(s.prefix, "empenho-desembolsos")
	var models []*schema.EmpenhoDesembolso
	total, err := s.get(&models, url, nil)
	return models, total, err
}
