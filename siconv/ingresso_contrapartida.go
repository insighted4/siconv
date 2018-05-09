package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *service) CreateIngressoContrapartida(ingressoContraPartida *schema.IngressoContrapartida) (string, error) {
	return s.dao.CreateIngressoContrapartida(ingressoContraPartida)
}

func (s *service) GetIngressoContrapartida(id string) (*schema.IngressoContrapartida, error) {
	return s.dao.GetIngressoContrapartida(id)
}

func (s *service) ListIngressoContrapartida(pagination *storage.Pagination) ([]*schema.IngressoContrapartida, int, error) {
	return s.dao.ListIngressoContrapartida(pagination)
}
