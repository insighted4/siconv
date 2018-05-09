package client

import (
	"path"

	"strconv"

	"fmt"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (s *Client) Create(model schema.Model) error {
	endpoint := newEndpoint(model)
	if endpoint == "" {
		return fmt.Errorf("unrecognized endpoint for %T", model)
	}

	url := path.Join(s.prefix, endpoint)
	location, err := s.post(model, url)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(path.Base(location))
	if err != nil {
		return fmt.Errorf("unable to convert location to id: %v", err)
	}

	model.SetID(id)
	return nil
}

func (s *Client) Get(model schema.Model) error {
	endpoint := newEndpoint(model)
	if endpoint == "" {
		return fmt.Errorf("unrecognized endpoint for %T", model)
	}

	url := path.Join(s.prefix, endpoint, strconv.Itoa(model.GetID()))
	_, err := s.get(&model, url, nil)
	return err
}

func (s *Client) List(models interface{}, pagination *storage.Pagination) (int, error) {
	endpoint := newEndpoint(models)
	if endpoint == "" {
		return 0, fmt.Errorf("unrecognized endpoint for %T", models)
	}

	url := path.Join(s.prefix, endpoint)
	total, err := s.get(&models, url, nil)
	return total, err
}

func newEndpoint(model interface{}) string {
	switch model.(type) {
	case *schema.Consorcio:
		return "consorcios"
	case []*schema.Consorcio:
		return "consorcios"

	case *schema.Convenio:
		return "convenios"
	case []*schema.Convenio:
		return "convenios"

	case *schema.Desembolso:
		return "desembolsos"
	case []*schema.Desembolso:
		return "desembolsos"

	case *schema.Emenda:
		return "emendas"
	case []*schema.Emenda:
		return "emendas"

	case *schema.Empenho:
		return "empenhos"
	case []*schema.Empenho:
		return "empenhos"

	case *schema.EmpenhoDesembolso:
		return "empenho"
	case []*schema.EmpenhoDesembolso:
		return "empenho"

	case *schema.EtapaCronoFisico:
		return "etapa-crono-fisico"
	case []*schema.EtapaCronoFisico:
		return "etapa-crono-fisico"

	case *schema.HistoricoSituacao:
		return "historico-situacoes"
	case []*schema.HistoricoSituacao:
		return "historico-situacoes"

	case *schema.IngressoContrapartida:
		return "ingresso-contrapartidas"
	case []*schema.IngressoContrapartida:
		return "ingresso-contrapartidas"

	case *schema.MetaCronoFisico:
		return "meta-crono-fisicos"
	case []*schema.MetaCronoFisico:
		return "meta-crono-fisicos"

	case *schema.OBTVConvenente:
		return "obtv-convenentes"
	case []*schema.OBTVConvenente:
		return "obtv-convenentes"

	case *schema.Pagamento:
		return "pagamentos"
	case []*schema.Pagamento:
		return "pagamentos"

	case *schema.PlanoAplicacaoDetalhado:
		return "plano-aplicacao-detalhados"
	case []*schema.PlanoAplicacaoDetalhado:
		return "plano-aplicacao-detalhados"

	case *schema.Programa:
		return "programas"
	case []*schema.Programa:
		return "programas"

	case *schema.ProgramaProposta:
		return "programa-propostas"
	case []*schema.ProgramaProposta:
		return "programa-propostas"

	case *schema.Proponente:
		return "proponentes"
	case []*schema.Proponente:
		return "proponentes"

	case *schema.Proposta:
		return "propostas"
	case []*schema.Proposta:
		return "propostas"

	case *schema.ProrrogaOficio:
		return "prorroga-oficios"
	case []*schema.ProrrogaOficio:
		return "prorroga-oficios"

	case *schema.TermoAditivo:
		return "termo-aditivos"
	case []*schema.TermoAditivo:
		return "termo-aditivos"
	}

	return ""
}
