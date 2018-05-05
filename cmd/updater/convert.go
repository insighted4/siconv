package main

import (
	"strings"
	"time"

	"strconv"

	"github.com/insighted4/siconv/schema"
)

func Boolean(value string) bool {
	return strings.ToUpper(value) == "SIM" || value == "1"
}

func Date(value string) *time.Time {
	date, err := time.Parse("02/01/2006", value)
	if err != nil {
		return nil
	}
	return &date
}

func Float64(value string) float64 {
	v := strings.Replace(value, ".", "", -1)
	v = strings.Replace(value, ",", ".", 1)

	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func Int(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func NewConsorcio(row map[string]string) *schema.Consorcio {
	return &schema.Consorcio{
		StorageModel:           schema.StorageModel{Reference: row["reference"]},
		ID_PROPOSTA:            row["id_proposta"],
		CNPJ_CONSORCIO:         row["cnpj_consorcio"],
		NOME_CONSORCIO:         row["nome_consorcio"],
		CODIGO_CNAE_PRIMARIO:   row["codigo_cnae_primario"],
		DESC_CNAE_PRIMARIO:     row["desc_cnae_primario"],
		CODIGO_CNAE_SECUNDARIO: row["codigo_cnae_secundario"],
		DESC_CNAE_SECUNDARIO:   row["desc_cnae_secundario"],
		CNPJ_PARTICIPANTE:      row["cnpj_participante"],
		NOME_PARTICIPANTE:      row["nome_participante"],
	}
}

func NewConvenio(row map[string]string) *schema.Convenio {
	return &schema.Convenio{
		StorageModel:              schema.StorageModel{Reference: row["reference"]},
		NR_CONVENIO:               row["nr_convenio"],
		ID_PROPOSTA:               row["id_proposta"],
		DIA_ASSIN_CONV:            Date(row["dia_assin_conv"]),
		SIT_CONVENIO:              row["sit_convenio"],
		SUBSITUACAO_CONV:          row["subsituacao_conv"],
		SITUACAO_PUBLICACAO:       row["situacao_publicacao"],
		INSTRUMENTO_ATIVO:         Boolean(row["instrumento_ativo"]),
		IND_OPERA_OBTV:            Boolean(row["ind_opera_obtv"]),
		NR_PROCESSO:               row["nr_processo"],
		UG_EMITENTE:               row["ug_emitente"],
		DIA_PUBL_CONV:             row["dia_publ_conv"],
		DIA_INIC_VIGENC_CONV:      Date(row["dia_inic_vigenc_conv"]),
		DIA_FIM_VIGENC_CONV:       Date(row["dia_fim_vigenc_conv"]),
		DIAS_PREST_CONTAS:         Date(row["dias_prest_contas"]),
		DIA_LIMITE_PREST_CONTAS:   Date(row["dia_limite_prest_contas"]),
		SITUACAO_CONTRATACAO:      row["situacao_contratacao"],
		IND_ASSINADO:              Boolean(row["ind_assinado"]),
		QTDE_CONVENIOS:            row["qtde_convenios"],
		QTD_TA:                    row["qtd_ta"],
		QTD_PRORROGA:              row["qtd_prorroga"],
		VL_GLOBAL_CONV:            Float64(row["vl_global_conv"]),
		VL_REPASSE_CONV:           Float64(row["vl_repasse_conv"]),
		VL_CONTRAPARTIDA_CONV:     Float64(row["vl_contrapartida_conv"]),
		VL_EMPENHADO_CONV:         Float64(row["vl_empenhado_conv"]),
		VL_DESEMBOLSADO_CONV:      Float64(row["vl_desembolsado_conv"]),
		VL_SALDO_REMAN_TESOURO:    Float64(row["vl_saldo_reman_tesouro"]),
		VL_SALDO_REMAN_CONVENENTE: Float64(row["vl_saldo_reman_convenente"]),
		VL_RENDIMENTO_APLICACAO:   Float64(row["vl_rendimento_aplicacao"]),
		VL_INGRESSO_CONTRAPARTIDA: Float64(row["vl_ingresso_contrapartida"]),
	}
}

func NewDesembolso(row map[string]string) *schema.Desembolso {
	return &schema.Desembolso{
		StorageModel:            schema.StorageModel{Reference: row["reference"]},
		NR_CONVENIO:             row["nr_convenio"],
		DT_ULT_DESEMBOLSO:       Date(row["dt_ult_desembolso"]),
		QTD_DIAS_SEM_DESEMBOLSO: Int(row["qtd_dias_sem_desembolso"]),
		ID_DESEMBOLSO:           row["id_desembolso"],
		DATA_DESEMBOLSO:         Date(row["data_desembolso"]),
		ANO_DESEMBOLSO:          Int(row["ano_desembolso"]),
		MES_DESEMBOLSO:          Int(row["mes_desembolso"]),
		NR_SIAFI:                row["nr_siafi"],
		VL_DESEMBOLSADO:         Float64(row["vl_desembolsado"]),
	}
}

func NewEmenda(row map[string]string) *schema.Emenda {
	return &schema.Emenda{
		StorageModel:                  schema.StorageModel{Reference: row["reference"]},
		ID_PROPOSTA:                   row["id_proposta"],
		QUALIF_PROPONENTE:             row["qualif_proponente"],
		COD_PROGRAMA_EMENDA:           row["cod_programa_emenda"],
		NR_EMENDA:                     row["nr_emenda"],
		NOME_PARLAMENTAR:              row["nome_parlamentar"],
		BENEFICIARIO_EMENDA:           row["beneficiario_emenda"],
		IND_IMPOSITIVO:                Boolean(row["ind_impositivo"]),
		TIPO_PARLAMENTAR:              row["tipo_parlamentar"],
		VALOR_REPASSE_PROPOSTA_EMENDA: Float64(row["valor_repasse_proposta_emenda"]),
		VALOR_REPASSE_EMENDA:          Float64(row["valor_repasse_emenda"]),
	}
}

func NewEmpenho(row map[string]string) *schema.Empenho {
	return &schema.Empenho{
		StorageModel:          schema.StorageModel{Reference: row["reference"]},
		ID_EMPENHO:            row["id_empenho"],
		NR_EMPENHO:            row["nr_empenho"],
		TIPO_NOTA:             row["tipo_nota"],
		DESC_TIPO_NOTA:        row["desc_tipo_nota"],
		DATA_EMISSAO:          Date(row["data_emissao"]),
		COD_SITUACAO_EMPENHO:  row["cod_situacao_empenho"],
		DESC_SITUACAO_EMPENHO: row["desc_situacao_empenho"],
		VALOR_EMPENHO:         Float64(row["valor_empenho"]),
		NR_CONVENIO:           row["nr_convenio"],
	}
}

func NewEmpenhoDesembolso(row map[string]string) *schema.EmpenhoDesembolso {
	return &schema.EmpenhoDesembolso{
		StorageModel:  schema.StorageModel{Reference: row["reference"]},
		ID_DESEMBOLSO: row["id_desembolso"],
		ID_EMPENHO:    row["id_empenho"],
		VALOR_GRUPO:   Float64(row["valor_grupo"]),
	}
}

func NewEtapaCronoFisico(row map[string]string) *schema.EtapaCronoFisico {
	return &schema.EtapaCronoFisico{
		StorageModel:           schema.StorageModel{Reference: row["reference"]},
		ID_META:                row["id_meta"],
		ID_ETAPA:               row["id_etapa"],
		NR_ETAPA:               row["nr_etapa"],
		DESC_ETAPA:             row["desc_etapa"],
		DATA_INICIO_ETAPA:      Date(row["data_inicio_etapa"]),
		DATA_FIM_ETAPA:         Date(row["data_fim_etapa"]),
		UF_ETAPA:               row["uf_etapa"],
		MUNICIPIO_ETAPA:        row["municipio_etapa"],
		ENDERECO_ETAPA:         row["endereco_etapa"],
		CEP_ETAPA:              row["cep_etapa"],
		QTD_ETAPA:              Int(row["qtd_etapa"]),
		UND_FORNECIMENTO_ETAPA: row["und_fornecimento_etapa"],
		VL_ETAPA:               Float64(row["vl_etapa"]),
	}
}

func NewHistoricoSituacao(row map[string]string) *schema.HistoricoSituacao {
	return &schema.HistoricoSituacao{
		StorageModel:       schema.StorageModel{Reference: row["reference"]},
		ID_PROPOSTA:        row["id_proposta"],
		NR_CONVENIO:        row["nr_convenio"],
		DIA_HISTORICO_SIT:  Date(row["dia_historico_sit"]),
		HISTORICO_SIT:      row["historico_sit"],
		DIAS_HISTORICO_SIT: Int(row["dias_historico_sit"]),
		COD_HISTORICO_SIT:  row["cod_historico_sit"],
	}
}

func NewIngressoContrapartida(row map[string]string) *schema.IngressoContrapartida {
	return &schema.IngressoContrapartida{
		StorageModel:              schema.StorageModel{Reference: row["reference"]},
		NR_CONVENIO:               row["nr_convenio"],
		DT_INGRESSO_CONTRAPARTIDA: Date(row["dt_ingresso_contrapartida"]),
		VL_INGRESSO_CONTRAPARTIDA: Float64(row["vl_ingresso_contrapartida"]),
	}
}

func NewMetaCronoFisico(row map[string]string) *schema.MetaCronoFisico {
	return &schema.MetaCronoFisico{
		StorageModel:          schema.StorageModel{Reference: row["reference"]},
		ID_META:               row["id_meta"],
		NR_CONVENIO:           row["nr_convenio"],
		COD_PROGRAMA:          row["cod_programa"],
		NOME_PROGRAMA:         row["nome_programa"],
		NR_META:               Int(row["nr_meta"]),
		TIPO_META:             row["tipo_meta"],
		DESC_META:             row["desc_meta"],
		DATA_INICIO_META:      Date(row["data_inicio_meta"]),
		DATA_FIM_META:         Date(row["data_fim_meta"]),
		UF_META:               row["uf_meta"],
		MUNICIPIO_META:        row["municipio_meta"],
		ENDERECO_META:         row["endereco_meta"],
		CEP_META:              row["cep_meta"],
		QTD_META:              Int(row["qtd_meta"]),
		UND_FORNECIMENTO_META: row["und_fornecimento_meta"],
		VL_META:               Float64(row["vl_meta"]),
	}
}

func NewOBTVConvenente(row map[string]string) *schema.OBTVConvenente {
	return &schema.OBTVConvenente{
		StorageModel:                 schema.StorageModel{Reference: row["reference"]},
		NR_MOV_FIN:                   row["nr_mov_fin"],
		IDENTIF_FAVORECIDO_OBTV_CONV: row["identif_favorecido_obtv_conv"],
		NM_FAVORECIDO_OBTV_CONV:      row["nm_favorecido_obtv_conv"],
		TP_AQUISICAO:                 row["tp_aquisicao"],
		VL_PAGO_OBTV_CONV:            Float64(row["vl_pago_obtv_conv"]),
	}
}

func NewPagamento(row map[string]string) *schema.Pagamento {
	return &schema.Pagamento{
		StorageModel:       schema.StorageModel{Reference: row["reference"]},
		NR_MOV_FIN:         row["nr_mov_fin"],
		NR_CONVENIO:        row["nr_convenio"],
		IDENTIF_FORNECEDOR: row["identif_fornecedor"],
		NOME_FORNECEDOR:    row["nome_fornecedor"],
		TP_MOV_FINANCEIRA:  row["tp_mov_financeira"],
		DATA_PAG:           Date(row["data_pag"]),
		NR_DL:              row["nr_dl"],
		DESC_DL:            row["desc_dl"],
		VL_PAGO:            Float64(row["vl_pago"]),
	}
}

func NewPlanoAplicacaoDetalhado(row map[string]string) *schema.PlanoAplicacaoDetalhado {
	return &schema.PlanoAplicacaoDetalhado{
		StorageModel:        schema.StorageModel{Reference: row["reference"]},
		ID_PROPOSTA:         row["id_proposta"],
		SIGLA:               row["sigla"],
		MUNICIPIO:           row["municipio"],
		NATUREZA_AQUISICAO:  Int(row["natureza_aquisicao"]),
		DESCRICAO_ITEM:      row["descricao_item"],
		CEP_ITEM:            row["cep_item"],
		ENDERECO_ITEM:       row["endereco_item"],
		TIPO_DESPESA_ITEM:   row["tipo_despesa_item"],
		NATUREZA_DESPESA:    row["natureza_despesa"],
		SIT_ITEM:            row["sit_item"],
		QTD_ITEM:            Int(row["qtd_item"]),
		VALOR_UNITARIO_ITEM: Float64(row["valor_unitario_item"]),
		VALOR_TOTAL_ITEM:    Float64(row["valor_total_item"]),
	}
}

func NewPrograma(row map[string]string) *schema.Programa {
	return &schema.Programa{
		StorageModel:               schema.StorageModel{Reference: row["reference"]},
		COD_ORGAO_SUP_PROGRAMA:     row["cod_orgao_sup_programa"],
		DESC_ORGAO_SUP_PROGRAMA:    row["desc_orgao_sup_programa"],
		ID_PROGRAMA:                row["id_programa"],
		COD_PROGRAMA:               row["cod_programa"],
		NOME_PROGRAMA:              row["nome_programa"],
		SIT_PROGRAMA:               row["sit_programa"],
		DATA_DISPONIBILIZACAO:      Date(row["data_disponibilizacao"]),
		ANO_DISPONIBILIZACAO:       row["ano_disponibilizacao"],
		DT_PROG_INI_RECEB_PROP:     Date(row["dt_prog_ini_receb_prop"]),
		DT_PROG_FIM_RECEB_PROP:     Date(row["dt_prog_fim_receb_prop"]),
		DT_PROG_INI_EMENDA_PAR:     Date(row["dt_prog_ini_emenda_par"]),
		DT_PROG_FIM_EMENDA_PAR:     Date(row["dt_prog_fim_emenda_par"]),
		DT_PROG_INI_BENEF_ESP:      Date(row["dt_prog_ini_benef_esp"]),
		DT_PROG_FIM_BENEF_ESP:      Date(row["dt_prog_fim_benef_esp"]),
		MODALIDADE_PROGRAMA:        row["modalidade_programa"],
		NATUREZA_JURIDICA_PROGRAMA: row["natureza_juridica_programa"],
		UF_PROGRAMA:                row["uf_programa"],
		ACAO_ORCAMENTARIA:          row["acao_orcamentaria"],
	}
}

func NewProgramaProposta(row map[string]string) *schema.ProgramaProposta {
	return &schema.ProgramaProposta{
		StorageModel: schema.StorageModel{Reference: row["reference"]},
		ID_PROGRAMA:  row["id_programa"],
		ID_PROPOSTA:  row["id_proposta"],
	}
}

func NewProponente(row map[string]string) *schema.Proponente {
	return &schema.Proponente{
		StorageModel:         schema.StorageModel{Reference: row["reference"]},
		IDENTIF_PROPONENTE:   row["identif_proponente"],
		NM_PROPONENTE:        row["nm_proponente"],
		MUNICIPIO_PROPONENTE: row["municipio_proponente"],
		UF_PROPONENTE:        row["uf_proponente"],
		ENDERECO_PROPONENTE:  row["endereco_proponente"],
		BAIRRO_PROPONENTE:    row["bairro_proponente"],
		CEP_PROPONENTE:       row["cep_proponente"],
		EMAIL_PROPONENTE:     row["email_proponente"],
		TELEFONE_PROPONENTE:  row["telefone_proponente"],
		FAX_PROPONENTE:       row["fax_proponente"],
	}
}

func NewProposta(row map[string]string) *schema.Proposta {
	return &schema.Proposta{
		StorageModel:               schema.StorageModel{Reference: row["reference"]},
		ID_PROPOSTA:                row["id_proposta"],
		UF_PROPONENTE:              row["uf_proponente"],
		MUNIC_PROPONENTE:           row["munic_proponente"],
		COD_MUNIC_IBGE:             row["cod_munic_ibge"],
		COD_ORGAO_SUP:              row["cod_orgao_sup"],
		DESC_ORGAO_SUP:             row["desc_orgao_sup"],
		NATUREZA_JURIDICA:          row["natureza_juridica"],
		NR_PROPOSTA:                row["nr_proposta"],
		DIA_PROPOSTA:               Date(row["dia_proposta"]),
		COD_ORGAO:                  row["cod_orgao"],
		DESC_ORGAO:                 row["desc_orgao"],
		MODALIDADE:                 row["modalidade"],
		IDENTIF_PROPONENTE:         row["identif_proponente"],
		NM_PROPONENTE:              row["nm_proponente"],
		CEP_PROPONENTE:             row["cep_proponente"],
		ENDERECO_PROPONENTE:        row["endereco_proponente"],
		BAIRRO_PROPONENTE:          row["bairro_proponente"],
		NM_BANCO:                   row["nm_banco"],
		SITUACAO_CONTA:             row["situacao_conta"],
		SITUACAO_PROJETO_BASICO:    row["situacao_projeto_basico"],
		SIT_PROPOSTA:               row["sit_proposta"],
		DIA_INIC_VIGENCIA_PROPOSTA: Date(row["dia_inic_vigencia_proposta"]),
		DIA_FIM_VIGENCIA_PROPOSTA:  Date(row["dia_fim_vigencia_proposta"]),
		OBJETO_PROPOSTA:            row["objeto_proposta"],
		VL_GLOBAL_PROP:             Float64(row["vl_global_prop"]),
		VL_REPASSE_PROP:            Float64(row["vl_repasse_prop"]),
		VL_CONTRAPARTIDA_PROP:      Float64(row["vl_contrapartida_prop"]),
	}
}

func NewProrrogaOficio(row map[string]string) *schema.ProrrogaOficio {
	return &schema.ProrrogaOficio{
		StorageModel:           schema.StorageModel{Reference: row["reference"]},
		NR_CONVENIO:            row["nr_convenio"],
		NR_PRORROGA:            row["nr_prorroga"],
		DT_INICIO_PRORROGA:     Date(row["dt_inicio_prorroga"]),
		DT_FIM_PRORROGA:        Date(row["dt_fim_prorroga"]),
		DIAS_PRORROGA:          Int(row["dias_prorroga"]),
		DT_ASSINATURA_PRORROGA: Date(row["dt_assinatura_prorroga"]),
		SIT_PRORROGA:           row["sit_prorroga"],
	}
}

func NewTermoAditivo(row map[string]string) *schema.TermoAditivo {
	return &schema.TermoAditivo{
		StorageModel:        schema.StorageModel{Reference: row["reference"]},
		NR_CONVENIO:         row["nr_convenio"],
		NUMERO_TA:           row["numero_ta"],
		TIPO_TA:             row["tipo_ta"],
		VL_GLOBAL_TA:        Float64(row["vl_global_ta"]),
		VL_REPASSE_TA:       Float64(row["vl_repasse_ta"]),
		VL_CONTRAPARTIDA_TA: Float64(row["vl_contrapartida_ta"]),
		DT_ASSINATURA_TA:    Date(row["dt_assinatura_ta"]),
		DT_INICIO_TA:        Date(row["dt_inicio_ta"]),
		DT_FIM_TA:           Date(row["dt_fim_ta"]),
		JUSTIFICATIVA_TA:    row["justificativa_ta"],
	}
}
