package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
)

func (s *Client) CreateEmpenhoDesembolso(empenhoDesembolso *schema.EmpenhoDesembolso) (string, error) {
	url := path.Join(s.prefix, "empenho-desembolsos")
	return s.post(empenhoDesembolso, url)
}
