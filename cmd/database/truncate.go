package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func truncate(filename string, storage *postgres, logger logrus.FieldLogger) error {
	var table string
	switch filename {
	case "siconv_consorcios.csv":
		table = "consorcios"
	case "siconv_convenio.csv":
		table = "convenios"
	case "siconv_desembolso.csv":
		table = "desembolsos"
	case "siconv_emenda.csv":
		table = "emendas"
	case "siconv_empenho.csv":
		table = "empenhos"
	case "siconv_empenho_desembolso.csv":
		table = "empenho_desembolsos"
	case "siconv_etapa_crono_fisico.csv":
		table = "etapa_crono_fisicos"
	case "siconv_historico_situacao.csv":
		table = "historico_situacaos"
	case "siconv_ingresso_contrapartida.csv":
		table = "ingresso_contrapartidas"
	case "siconv_meta_crono_fisico.csv":
		table = "meta_crono_fisicos"
	case "siconv_obtv_convenente.csv":
		table = "obtv_convenentes"
	case "siconv_pagamento.csv":
		table = "pagamentos"
	case "siconv_plano_aplicacao_detalhado.csv":
		table = "plano_aplicacao_detalhados"
	case "siconv_programa.csv":
		table = "programas"
	case "siconv_programa_proposta.csv":
		table = "programa_propostas"
	case "siconv_proponentes.csv":
		table = "proponentes"
	case "siconv_proposta.csv":
		table = "propostas"
	case "siconv_prorroga_oficio.csv":
		table = "prorroga_oficios"
	case "siconv_termo_aditivo.csv":
		table = "termo_aditivos"
	default:
		return fmt.Errorf("unrecognized filename")
	}

	logger.Warnf("Truncating table %s", table)
	_, err := storage.db.ExecOne("TRUNCATE TABLE " + table)
	return err
}
