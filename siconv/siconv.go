package siconv

import (
	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateConsorcio(consorcio *schema.Consorcio) (id string, err error)
	GetConsorcio(id string) (*schema.Consorcio, error)
	ListConsorcio(pagination *storage.Pagination) ([]*schema.Consorcio, int, error)

	CreateConvenio(convenio *schema.Convenio) (string, error)
	GetConvenio(id string) (*schema.Convenio, error)
	ListConvenio(pagination *storage.Pagination) ([]*schema.Convenio, int, error)

	CreateDesembolso(desembolso *schema.Desembolso) (string, error)
	GetDesembolso(id string) (*schema.Desembolso, error)
	ListDesembolso(pagination *storage.Pagination) ([]*schema.Desembolso, int, error)

	CreateEmenda(emenda *schema.Emenda) (string, error)
	GetEmenda(id string) (*schema.Emenda, error)
	ListEmenda(pagination *storage.Pagination) ([]*schema.Emenda, int, error)

	CreateEmpenho(empenho *schema.Empenho) (string, error)
	GetEmpenho(id string) (*schema.Empenho, error)
	ListEmpenho(pagination *storage.Pagination) ([]*schema.Empenho, int, error)

	CreateEmpenhoDesembolso(empanhoDesembolso *schema.EmpenhoDesembolso) (string, error)
	ListEmpenhoDesembolso(pagination *storage.Pagination) ([]*schema.EmpenhoDesembolso, int, error)

	CreateEtapaCronoFisico(etapaCronoFisico *schema.EtapaCronoFisico) (string, error)
	GetEtapaCronoFisico(id string) (*schema.EtapaCronoFisico, error)
	ListEtapaCronoFisico(pagination *storage.Pagination) ([]*schema.EtapaCronoFisico, int, error)

	CreateHistoricoSituacao(historicoSituacao *schema.HistoricoSituacao) (string, error)
	GetHistoricoSituacao(id string) (*schema.HistoricoSituacao, error)
	ListHistoricoSituacao(pagination *storage.Pagination) ([]*schema.HistoricoSituacao, int, error)

	CreateIngressoContrapartida(ingressoContrapartida *schema.IngressoContrapartida) (string, error)
	GetIngressoContrapartida(id string) (*schema.IngressoContrapartida, error)
	ListIngressoContrapartida(pagination *storage.Pagination) ([]*schema.IngressoContrapartida, int, error)

	CreateMetaCronoFisico(metaCronoFisico *schema.MetaCronoFisico) (string, error)
	GetMetaCronoFisico(id string) (*schema.MetaCronoFisico, error)
	ListMetaCronoFisico(pagination *storage.Pagination) ([]*schema.MetaCronoFisico, int, error)

	CreateOBTVConvenente(obtvConvenente *schema.OBTVConvenente) (string, error)
	GetOBTVConvenente(id string) (*schema.OBTVConvenente, error)
	ListOBTVConvenente(pagination *storage.Pagination) ([]*schema.OBTVConvenente, int, error)

	CreatePagamento(pagamento *schema.Pagamento) (string, error)
	GetPagamento(id string) (*schema.Pagamento, error)
	ListPagamento(pagination *storage.Pagination) ([]*schema.Pagamento, int, error)

	CreatePlanoAplicacaoDetalhado(planoAplicacaoDetalhado *schema.PlanoAplicacaoDetalhado) (string, error)
	GetPlanoAplicacaoDetalhado(id string) (*schema.PlanoAplicacaoDetalhado, error)
	ListPlanoAplicacaoDetalhado(pagination *storage.Pagination) ([]*schema.PlanoAplicacaoDetalhado, int, error)

	CreatePrograma(programa *schema.Programa) (string, error)
	GetPrograma(id string) (*schema.Programa, error)
	ListPrograma(idPrograma string, pagination *storage.Pagination) ([]*schema.Programa, int, error)

	CreateProgramaProposta(programaProposta *schema.ProgramaProposta) (string, error)
	ListProgramaProposta(idPrograma string, pagination *storage.Pagination) ([]*schema.Proposta, int, error)

	CreateProponente(proponente *schema.Proponente) (string, error)
	GetProponente(id string) (*schema.Proponente, error)
	ListProponente(pagination *storage.Pagination) ([]*schema.Proponente, int, error)

	CreateProposta(proposta *schema.Proposta) (string, error)
	GetProposta(id string) (*schema.Proposta, error)
	ListProposta(pagination *storage.Pagination) ([]*schema.Proposta, int, error)

	CreateProrrogaOficio(prorrogaOficio *schema.ProrrogaOficio) (string, error)
	GetProrrogaOficio(id string) (*schema.ProrrogaOficio, error)
	ListProrrogaOficio(pagination *storage.Pagination) ([]*schema.ProrrogaOficio, int, error)

	CreateTermoAditivo(termoAditivo *schema.TermoAditivo) (string, error)
	GetTermoAditivo(id string) (*schema.TermoAditivo, error)
	ListTermoAditivo(pagination *storage.Pagination) ([]*schema.TermoAditivo, int, error)
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
