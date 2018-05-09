package client

import (
	"path"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *Client) CreateProgramaProposta(programaProposta *schema.ProgramaProposta) (string, error) {
	url := path.Join(s.prefix, "programa-propostas")
	return s.post(programaProposta, url)
}

func (s *Client) ListProgramaProposta(idPrograma string, pagination *storage.Pagination) ([]*schema.Proposta, int, error) {
	url := path.Join(s.prefix, "programas", idPrograma, "propostas")
	var models []*schema.Proposta
	total, err := s.get(&models, url, nil)
	return models, total, err
}
