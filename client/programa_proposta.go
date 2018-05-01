package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (s *Client) CreateProgramaProposta(programaProposta *schema.ProgramaProposta) (string, error) {
	url := path.Join(s.prefix, "programa-propostas")
	return s.post(programaProposta, url)
}

func (s *Client) ListProgramaProposta(idPrograma string, pagination *siconv.Pagination) ([]*schema.Proposta, int, error) {
	url := path.Join(s.prefix, "programas", idPrograma, "propostas")
	var models []*schema.Proposta
	total, err := s.get(&models, url, nil)
	return models, total, err
}
