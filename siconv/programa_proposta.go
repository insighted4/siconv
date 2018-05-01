package siconv

import (
	"github.com/insighted4/siconv/schema"
)

func (s *service) CreateProgramaProposta(programaProposta *schema.ProgramaProposta) (string, error) {
	return s.dao.CreateProgramaProposta(programaProposta)
}

func (s *service) ListProgramaProposta(idPrograma string, pagination *Pagination) ([]*schema.Proposta, int, error) {
	return s.dao.ListProgramaProposta(idPrograma, pagination)
}
