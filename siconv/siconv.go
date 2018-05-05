package siconv

import (
	"errors"
	"regexp"

	"github.com/insighted4/siconv/schema"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateConsorcio(consorcio *schema.Consorcio) (id string, err error)
	GetConsorcio(id string) (*schema.Consorcio, error)
	ListConsorcio(pagination *Pagination) ([]*schema.Consorcio, int, error)

	CreateConvenio(convenio *schema.Convenio) (string, error)
	GetConvenio(id string) (*schema.Convenio, error)
	ListConvenio(pagination *Pagination) ([]*schema.Convenio, int, error)

	CreateDesembolso(desembolso *schema.Desembolso) (string, error)
	GetDesembolso(id string) (*schema.Desembolso, error)
	ListDesembolso(pagination *Pagination) ([]*schema.Desembolso, int, error)

	CreateEmenda(emenda *schema.Emenda) (string, error)
	GetEmenda(id string) (*schema.Emenda, error)
	ListEmenda(pagination *Pagination) ([]*schema.Emenda, int, error)

	CreateEmpenho(empenho *schema.Empenho) (string, error)
	GetEmpenho(id string) (*schema.Empenho, error)
	ListEmpenho(pagination *Pagination) ([]*schema.Empenho, int, error)

	CreateEmpenhoDesembolso(empanhoDesembolso *schema.EmpenhoDesembolso) (string, error)
	ListEmpenhoDesembolso(pagination *Pagination) ([]*schema.EmpenhoDesembolso, int, error)

	CreateEtapaCronoFisico(etapaCronoFisico *schema.EtapaCronoFisico) (string, error)
	GetEtapaCronoFisico(id string) (*schema.EtapaCronoFisico, error)
	ListEtapaCronoFisico(pagination *Pagination) ([]*schema.EtapaCronoFisico, int, error)

	CreateHistoricoSituacao(historicoSituacao *schema.HistoricoSituacao) (string, error)
	GetHistoricoSituacao(id string) (*schema.HistoricoSituacao, error)
	ListHistoricoSituacao(pagination *Pagination) ([]*schema.HistoricoSituacao, int, error)

	CreateIngressoContrapartida(ingressoContrapartida *schema.IngressoContrapartida) (string, error)
	GetIngressoContrapartida(id string) (*schema.IngressoContrapartida, error)
	ListIngressoContrapartida(pagination *Pagination) ([]*schema.IngressoContrapartida, int, error)

	CreateMetaCronoFisico(metaCronoFisico *schema.MetaCronoFisico) (string, error)
	GetMetaCronoFisico(id string) (*schema.MetaCronoFisico, error)
	ListMetaCronoFisico(pagination *Pagination) ([]*schema.MetaCronoFisico, int, error)

	CreateOBTVConvenente(obtvConvenente *schema.OBTVConvenente) (string, error)
	GetOBTVConvenente(id string) (*schema.OBTVConvenente, error)
	ListOBTVConvenente(pagination *Pagination) ([]*schema.OBTVConvenente, int, error)

	CreatePagamento(pagamento *schema.Pagamento) (string, error)
	GetPagamento(id string) (*schema.Pagamento, error)
	ListPagamento(pagination *Pagination) ([]*schema.Pagamento, int, error)

	CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado *schema.PlanoAplicacaoDetalhado) (string, error)
	GetPlanoAplicacaoDetalhado(id string) (*schema.PlanoAplicacaoDetalhado, error)
	ListPlanoAplicacaoDetalhado(pagination *Pagination) ([]*schema.PlanoAplicacaoDetalhado, int, error)

	CreatePrograma(programa *schema.Programa) (string, error)
	GetPrograma(id string) (*schema.Programa, error)
	ListPrograma(idPrograma string, pagination *Pagination) ([]*schema.Programa, int, error)

	CreateProgramaProposta(programaProposta *schema.ProgramaProposta) (string, error)
	ListProgramaProposta(idPrograma string, pagination *Pagination) ([]*schema.Proposta, int, error)

	CreateProponente(proponente *schema.Proponente) (string, error)
	GetProponente(id string) (*schema.Proponente, error)
	ListProponente(pagination *Pagination) ([]*schema.Proponente, int, error)

	CreateProposta(proposta *schema.Proposta) (string, error)
	GetProposta(id string) (*schema.Proposta, error)
	ListProposta(pagination *Pagination) ([]*schema.Proposta, int, error)

	CreateProrrogaOficio(prorrogaOficio *schema.ProrrogaOficio) (string, error)
	GetProrrogaOficio(id string) (*schema.ProrrogaOficio, error)
	ListProrrogaOficio(pagination *Pagination) ([]*schema.ProrrogaOficio, int, error)

	CreateTermoAditivo(termoAditivo *schema.TermoAditivo) (string, error)
	GetTermoAditivo(id string) (*schema.TermoAditivo, error)
	ListTermoAditivo(pagination *Pagination) ([]*schema.TermoAditivo, int, error)
}

type service struct {
	dao    Service
	logger logrus.FieldLogger
}

func New(dao Service, logger logrus.FieldLogger) Service {
	return &service{
		dao:    dao,
		logger: logger.WithField("component", "siconv"),
	}
}

var (
	// Limit default
	Limit = 100

	// ErrNotFound is the error returned by storage if a resource cannot be found.
	ErrNotFound = errors.New("not found")

	// ErrAlreadyExists is the error returned by storage if a resource ID is taken during a create.
	ErrAlreadyExists = errors.New("ID already exists")

	// ErrInvalidUUID is the error returned by storage if ID is not valid UUID.
	ErrInvalidUUID = errors.New("invalid UUID")
)

type Pagination struct {
	Limit  int
	Offset int
}

// NewPagination is passed as a parameter to limit the total of rows.
func NewPagination(perPage, page int) *Pagination {
	return &Pagination{
		Limit:  perPage,
		Offset: page * perPage,
	}
}

// IsValidUUID checks if a given string is a valid UUID v4.
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	return r.MatchString(uuid)
}

// IsValidUUID checks if a given string is a valid UUID v4.
func IsValidUUIDV4(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
