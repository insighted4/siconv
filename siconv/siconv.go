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

	CreateConvenio(convenio *schema.Convenio) (string, error)
	GetConvenio(id string) (*schema.Convenio, error)

	CreateDesembolso(desembolso *schema.Desembolso) (string, error)
	GetDesembolso(id string) (*schema.Desembolso, error)

	CreateEmenda(emenda *schema.Emenda) (string, error)
	GetEmenda(id string) (*schema.Emenda, error)

	CreateEmpenho(empenho *schema.Empenho) (string, error)
	GetEmpenho(id string) (*schema.Empenho, error)

	CreateEmpenhoDesembolso(empanhoDesembolso *schema.EmpenhoDesembolso) (string, error)

	CreateEtapaCronoFisico(etapaCronoFisico *schema.EtapaCronoFisico) (string, error)
	GetEtapaCronoFisico(id string) (*schema.EtapaCronoFisico, error)

	CreateHistoricoSituacao(historicoSituacao *schema.HistoricoSituacao) (string, error)
	GetHistoricoSituacao(id string) (*schema.HistoricoSituacao, error)

	CreateIngressoContrapartida(ingressoContrapartida *schema.IngressoContrapartida) (string, error)
	GetIngressoContrapartida(id string) (*schema.IngressoContrapartida, error)

	CreateMetaCronoFisico(metaCronoFisico *schema.MetaCronoFisico) (string, error)
	GetMetaCronoFisico(id string) (*schema.MetaCronoFisico, error)

	CreateOBTVConvenente(obtvConvenente *schema.OBTVConvenente) (string, error)
	GetOBTVConvenente(id string) (*schema.OBTVConvenente, error)

	CreatePagamento(pagamento *schema.Pagamento) (string, error)
	GetPagamento(id string) (*schema.Pagamento, error)

	CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado *schema.PlanoAplicacaoDetalhado) (string, error)
	GetPlanoAplicacaoDetalhado(id string) (*schema.PlanoAplicacaoDetalhado, error)

	CreatePrograma(programa *schema.Programa) (string, error)
	GetPrograma(id string) (*schema.Programa, error)
	ListPrograma(idPrograma string, pagination *Pagination) ([]*schema.Programa, int, error)

	CreateProgramaProposta(programaProposta *schema.ProgramaProposta) (string, error)
	ListProgramaProposta(idPrograma string, pagination *Pagination) ([]*schema.Proposta, int, error)

	CreateProponente(proponente *schema.Proponente) (string, error)
	GetProponente(id string) (*schema.Proponente, error)

	CreateProposta(proposta *schema.Proposta) (string, error)
	GetProposta(id string) (*schema.Proposta, error)

	CreateProrrogaOficio(prorrogaOficio *schema.ProrrogaOficio) (string, error)
	GetProrrogaOficio(id string) (*schema.ProrrogaOficio, error)

	CreateTermoAditivo(termoAditivo *schema.TermoAditivo) (string, error)
	GetTermoAditivo(id string) (*schema.TermoAditivo, error)
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
