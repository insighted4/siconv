package schema

import "time"

type Model interface {
	GetID() string
}

type StorageModel struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *StorageModel) GetID() string {
	return m.ID
}

type Consorcio struct {
	StorageModel
	ID_PROPOSTA            string `json:"id_proposta"`
	CNPJ_CONSORCIO         string `json:"cnpj_consorcio"`
	NOME_CONSORCIO         string `json:"nome_consorcio"`
	CODIGO_CNAE_PRIMARIO   string `json:"codigo_cnae_primario"`
	DESC_CNAE_PRIMARIO     string `json:"desc_cnae_primario"`
	CODIGO_CNAE_SECUNDARIO string `json:"codigo_cnae_secundario"`
	DESC_CNAE_SECUNDARIO   string `json:"desc_cnae_secundario"`
	CNPJ_PARTICIPANTE      string `json:"cnpj_participante"`
	NOME_PARTICIPANTE      string `json:"nome_participante"`
}

type Convenio struct {
	StorageModel
	NR_CONVENIO               string     `json:"nr_convenio"`
	ID_PROPOSTA               string     `json:"id_proposta"`
	DIA_ASSIN_CONV            *time.Time `json:"dia_assin_conv"`
	SIT_CONVENIO              string     `json:"sit_convenio"`
	SUBSITUACAO_CONV          string     `json:"subsituacao_conv"`
	SITUACAO_PUBLICACAO       string     `json:"situacao_publicacao"`
	INSTRUMENTO_ATIVO         bool       `json:"instrumento_ativo"`
	IND_OPERA_OBTV            bool       `json:"ind_opera_obtv"`
	NR_PROCESSO               string     `json:"nr_processo"`
	UG_EMITENTE               string     `json:"ug_emitente"`
	DIA_PUBL_CONV             string     `json:"dia_publ_conv"`
	DIA_INIC_VIGENC_CONV      *time.Time `json:"dia_inic_vigenc_conv"`
	DIA_FIM_VIGENC_CONV       *time.Time `json:"dia_fim_vigenc_conv"`
	DIAS_PREST_CONTAS         *time.Time `json:"dias_prest_contas"`
	DIA_LIMITE_PREST_CONTAS   *time.Time `json:"dia_limite_prest_contas"`
	SITUACAO_CONTRATACAO      string     `json:"situacao_contratacao"`
	IND_ASSINADO              bool       `json:"ind_assinado"`
	QTDE_CONVENIOS            string     `json:"qtde_convenios"`
	QTD_TA                    string     `json:"qtd_ta"`
	QTD_PRORROGA              string     `json:"qtd_prorroga"`
	VL_GLOBAL_CONV            float64    `json:"vl_global_conv"`
	VL_REPASSE_CONV           float64    `json:"vl_repasse_conv"`
	VL_CONTRAPARTIDA_CONV     float64    `json:"vl_contrapartida_conv"`
	VL_EMPENHADO_CONV         float64    `json:"vl_empenhado_conv"`
	VL_DESEMBOLSADO_CONV      float64    `json:"vl_desembolsado_conv"`
	VL_SALDO_REMAN_TESOURO    float64    `json:"vl_saldo_reman_tesouro"`
	VL_SALDO_REMAN_CONVENENTE float64    `json:"vl_saldo_reman_convenente"`
	VL_RENDIMENTO_APLICACAO   float64    `json:"vl_rendimento_aplicacao"`
	VL_INGRESSO_CONTRAPARTIDA float64    `json:"vl_ingresso_contrapartida"`
}

type Desembolso struct {
	StorageModel
	NR_CONVENIO             string     `json:"nr_convenio"`
	DT_ULT_DESEMBOLSO       *time.Time `json:"dt_ult_desembolso"`
	QTD_DIAS_SEM_DESEMBOLSO int        `json:"qtd_dias_sem_desembolso"`
	ID_DESEMBOLSO           string     `json:"id_desembolso"`
	DATA_DESEMBOLSO         *time.Time `json:"data_desembolso"`
	ANO_DESEMBOLSO          int        `json:"ano_desembolso"`
	MES_DESEMBOLSO          int        `json:"mes_desembolso"`
	NR_SIAFI                string     `json:"nr_siafi"`
	VL_DESEMBOLSADO         float64    `json:"vl_desembolsado"`
}

type Emenda struct {
	StorageModel
	ID_PROPOSTA                   string  `json:"id_proposta"`
	QUALIF_PROPONENTE             string  `json:"qualif_proponente"`
	COD_PROGRAMA_EMENDA           string  `json:"cod_programa_emenda"`
	NR_EMENDA                     string  `json:"nr_emenda"`
	NOME_PARLAMENTAR              string  `json:"nome_parlamentar"`
	BENEFICIARIO_EMENDA           string  `json:"beneficiario_emenda"`
	IND_IMPOSITIVO                bool    `json:"ind_impositivo"`
	TIPO_PARLAMENTAR              string  `json:"tipo_parlamentar"`
	VALOR_REPASSE_PROPOSTA_EMENDA float64 `json:"valor_repasse_proposta_emenda"`
	VALOR_REPASSE_EMENDA          float64 `json:"valor_repasse_emenda"`
}

type Empenho struct {
	StorageModel
	ID_EMPENHO            string     `json:"id_empenho"`
	NR_EMPENHO            string     `json:"nr_empenho"`
	TIPO_NOTA             string     `json:"tipo_nota"`
	DESC_TIPO_NOTA        string     `json:"desc_tipo_nota"`
	DATA_EMISSAO          *time.Time `json:"data_emissao"`
	COD_SITUACAO_EMPENHO  string     `json:"cod_situacao_empenho"`
	DESC_SITUACAO_EMPENHO string     `json:"desc_situacao_empenho"`
	VALOR_EMPENHO         float64    `json:"valor_empenho"`
	NR_CONVENIO           string     `json:"nr_convenio"`
}

type EmpenhoDesembolso struct {
	StorageModel
	ID_DESEMBOLSO string  `json:"id_desembolso"`
	ID_EMPENHO    string  `json:"id_empenho"`
	VALOR_GRUPO   float64 `json:"valor_grupo"`
}

type EtapaCronoFisico struct {
	StorageModel
	ID_META                string     `json:"id_meta"`
	ID_ETAPA               string     `json:"id_etapa"`
	NR_ETAPA               string     `json:"nr_etapa"`
	DESC_ETAPA             string     `json:"desc_etapa"`
	DATA_INICIO_ETAPA      *time.Time `json:"data_inicio_etapa"`
	DATA_FIM_ETAPA         *time.Time `json:"data_fim_etapa"`
	UF_ETAPA               string     `json:"uf_etapa"`
	MUNICIPIO_ETAPA        string     `json:"municipio_etapa"`
	ENDERECO_ETAPA         string     `json:"endereco_etapa"`
	CEP_ETAPA              string     `json:"cep_etapa"`
	QTD_ETAPA              int        `json:"qtd_etapa"`
	UND_FORNECIMENTO_ETAPA string     `json:"und_fornecimento_etapa"`
	VL_ETAPA               float64    `json:"vl_etapa"`
}

type HistoricoSituacao struct {
	StorageModel
	ID_PROPOSTA        string     `json:"id_proposta"`
	NR_CONVENIO        string     `json:"nr_convenio"`
	DIA_HISTORICO_SIT  *time.Time `json:"dia_historico_sit"`
	HISTORICO_SIT      string     `json:"historico_sit"`
	DIAS_HISTORICO_SIT int        `json:"dias_historico_sit"`
	COD_HISTORICO_SIT  string     `json:"cod_historico_sit"`

	TableName struct{} `sql:"historico_situacoes"`
}

type IngressoContrapartida struct {
	StorageModel
	NR_CONVENIO               string     `json:"nr_convenio"`
	DT_INGRESSO_CONTRAPARTIDA *time.Time `json:"dt_ingresso_contrapartida"`
	VL_INGRESSO_CONTRAPARTIDA float64    `json:"vl_ingresso_contrapartida"`
}

type MetaCronoFisico struct {
	StorageModel
	ID_META               string     `json:"id_meta"`
	NR_CONVENIO           string     `json:"nr_convenio"`
	COD_PROGRAMA          string     `json:"cod_programa"`
	NOME_PROGRAMA         string     `json:"nome_programa"`
	NR_META               int        `json:"nr_meta"`
	TIPO_META             string     `json:"tipo_meta"`
	DESC_META             string     `json:"desc_meta"`
	DATA_INICIO_META      *time.Time `json:"data_inicio_meta"`
	DATA_FIM_META         *time.Time `json:"data_fim_meta"`
	UF_META               string     `json:"uf_meta"`
	MUNICIPIO_META        string     `json:"municipio_meta"`
	ENDERECO_META         string     `json:"endereco_meta"`
	CEP_META              string     `json:"cep_meta"`
	QTD_META              int        `json:"qtd_meta"`
	UND_FORNECIMENTO_META string     `json:"und_fornecimento_meta"`
	VL_META               float64    `json:"vl_meta"`
}

type OBTVConvenente struct {
	StorageModel
	NR_MOV_FIN                   string  `json:"nr_mov_fin"`
	IDENTIF_FAVORECIDO_OBTV_CONV string  `json:"identif_favorecido_obtv_conv"`
	NM_FAVORECIDO_OBTV_CONV      string  `json:"nm_favorecido_obtv_conv"`
	TP_AQUISICAO                 string  `json:"tp_aquisicao"`
	VL_PAGO_OBTV_CONV            float64 `json:"vl_pago_obtv_conv"`
}

type Pagamento struct {
	StorageModel
	NR_MOV_FIN         string     `json:"nr_mov_fin"`
	NR_CONVENIO        string     `json:"nr_convenio"`
	IDENTIF_FORNECEDOR string     `json:"identif_fornecedor"`
	NOME_FORNECEDOR    string     `json:"nome_fornecedor"`
	TP_MOV_FINANCEIRA  string     `json:"tp_mov_financeira"`
	DATA_PAG           *time.Time `json:"data_pag"`
	NR_DL              string     `json:"nr_dl"`
	DESC_DL            string     `json:"desc_dl"`
	VL_PAGO            float64    `json:"vl_pago"`
}

type PlanoAplicacaoDetalhado struct {
	StorageModel
	ID_PROPOSTA         string  `json:"id_proposta"`
	SIGLA               string  `json:"sigla"`
	MUNICIPIO           string  `json:"municipio"`
	NATUREZA_AQUISICAO  int     `json:"natureza_aquisicao"`
	DESCRICAO_ITEM      string  `json:"descricao_item"`
	CEP_ITEM            string  `json:"cep_item"`
	ENDERECO_ITEM       string  `json:"endereco_item"`
	TIPO_DESPESA_ITEM   string  `json:"tipo_despesa_item"`
	NATUREZA_DESPESA    string  `json:"natureza_despesa"`
	SIT_ITEM            string  `json:"sit_item"`
	QTD_ITEM            int     `json:"qtd_item"`
	VALOR_UNITARIO_ITEM float64 `json:"valor_unitario_item"`
	VALOR_TOTAL_ITEM    float64 `json:"valor_total_item"`
}

type Programa struct {
	StorageModel
	COD_ORGAO_SUP_PROGRAMA     string     `json:"cod_orgao_sup_programa"`
	DESC_ORGAO_SUP_PROGRAMA    string     `json:"desc_orgao_sup_programa"`
	ID_PROGRAMA                string     `json:"id_programa"`
	COD_PROGRAMA               string     `json:"cod_programa"`
	NOME_PROGRAMA              string     `json:"nome_programa"`
	SIT_PROGRAMA               string     `json:"sit_programa"`
	DATA_DISPONIBILIZACAO      *time.Time `json:"data_disponibilizacao"`
	ANO_DISPONIBILIZACAO       string     `json:"ano_disponibilizacao"`
	DT_PROG_INI_RECEB_PROP     *time.Time `json:"dt_prog_ini_receb_prop"`
	DT_PROG_FIM_RECEB_PROP     *time.Time `json:"dt_prog_fim_receb_prop"`
	DT_PROG_INI_EMENDA_PAR     *time.Time `json:"dt_prog_ini_emenda_par"`
	DT_PROG_FIM_EMENDA_PAR     *time.Time `json:"dt_prog_fim_emenda_par"`
	DT_PROG_INI_BENEF_ESP      *time.Time `json:"dt_prog_ini_benef_esp"`
	DT_PROG_FIM_BENEF_ESP      *time.Time `json:"dt_prog_fim_benef_esp"`
	MODALIDADE_PROGRAMA        string     `json:"modalidade_programa"`
	NATUREZA_JURIDICA_PROGRAMA string     `json:"natureza_juridica_programa"`
	UF_PROGRAMA                string     `json:"uf_programa"`
	ACAO_ORCAMENTARIA          string     `json:"acao_orcamentaria"`
}

type ProgramaProposta struct {
	StorageModel
	ID_PROGRAMA string `json:"id_programa"`
	ID_PROPOSTA string `json:"id_proposta"`

	TableName struct{} `sql:"programa_propostas"`
}

type Proponente struct {
	StorageModel
	IDENTIF_PROPONENTE   string `json:"identif_proponente"`
	NM_PROPONENTE        string `json:"nm_proponente"`
	MUNICIPIO_PROPONENTE string `json:"municipio_proponente"`
	UF_PROPONENTE        string `json:"uf_proponente"`
	ENDERECO_PROPONENTE  string `json:"endereco_proponente"`
	BAIRRO_PROPONENTE    string `json:"bairro_proponente"`
	CEP_PROPONENTE       string `json:"cep_proponente"`
	EMAIL_PROPONENTE     string `json:"email_proponente"`
	TELEFONE_PROPONENTE  string `json:"telefone_proponente"`
	FAX_PROPONENTE       string `json:"fax_proponente"`
}

type Proposta struct {
	StorageModel
	ID_PROPOSTA                string     `json:"id_proposta"`
	UF_PROPONENTE              string     `json:"uf_proponente"`
	MUNIC_PROPONENTE           string     `json:"munic_proponente"`
	COD_MUNIC_IBGE             string     `json:"cod_munic_ibge"`
	COD_ORGAO_SUP              string     `json:"cod_orgao_sup"`
	DESC_ORGAO_SUP             string     `json:"desc_orgao_sup"`
	NATUREZA_JURIDICA          string     `json:"natureza_juridica"`
	NR_PROPOSTA                string     `json:"nr_proposta"`
	DIA_PROPOSTA               *time.Time `json:"dia_proposta"`
	COD_ORGAO                  string     `json:"cod_orgao"`
	DESC_ORGAO                 string     `json:"desc_orgao"`
	MODALIDADE                 string     `json:"modalidade"`
	IDENTIF_PROPONENTE         string     `json:"identif_proponente"`
	NM_PROPONENTE              string     `json:"nm_proponente"`
	CEP_PROPONENTE             string     `json:"cep_proponente"`
	ENDERECO_PROPONENTE        string     `json:"endereco_proponente"`
	BAIRRO_PROPONENTE          string     `json:"bairro_proponente"`
	NM_BANCO                   string     `json:"nm_banco"`
	SITUACAO_CONTA             string     `json:"situacao_conta"`
	SITUACAO_PROJETO_BASICO    string     `json:"situacao_projeto_basico"`
	SIT_PROPOSTA               string     `json:"sit_proposta"`
	DIA_INIC_VIGENCIA_PROPOSTA *time.Time `json:"dia_inic_vigencia_proposta"`
	DIA_FIM_VIGENCIA_PROPOSTA  *time.Time `json:"dia_fim_vigencia_proposta"`
	OBJETO_PROPOSTA            string     `json:"objeto_proposta"`
	VL_GLOBAL_PROP             float64    `json:"vl_global_prop"`
	VL_REPASSE_PROP            float64    `json:"vl_repasse_prop"`
	VL_CONTRAPARTIDA_PROP      float64    `json:"vl_contrapartida_prop"`

	Convenios []*Convenio `json:"convenios" sql:"-"`

	TableName struct{} `sql:"propostas"`
}

type ProrrogaOficio struct {
	StorageModel
	NR_CONVENIO            string     `json:"nr_convenio"`
	NR_PRORROGA            string     `json:"nr_prorroga"`
	DT_INICIO_PRORROGA     *time.Time `json:"dt_inicio_prorroga"`
	DT_FIM_PRORROGA        *time.Time `json:"dt_fim_prorroga"`
	DIAS_PRORROGA          int        `json:"dias_prorroga"`
	DT_ASSINATURA_PRORROGA *time.Time `json:"dt_assinatura_prorroga"`
	SIT_PRORROGA           string     `json:"sit_prorroga"`
}

type TermoAditivo struct {
	StorageModel
	NR_CONVENIO         string     `json:"nr_convenio"`
	NUMERO_TA           string     `json:"numero_ta"`
	TIPO_TA             string     `json:"tipo_ta"`
	VL_GLOBAL_TA        float64    `json:"vl_global_ta"`
	VL_REPASSE_TA       float64    `json:"vl_repasse_ta"`
	VL_CONTRAPARTIDA_TA float64    `json:"vl_contrapartida_ta"`
	DT_ASSINATURA_TA    *time.Time `json:"dt_assinatura_ta"`
	DT_INICIO_TA        *time.Time `json:"dt_inicio_ta"`
	DT_FIM_TA           *time.Time `json:"dt_fim_ta"`
	JUSTIFICATIVA_TA    string     `json:"justificativa_ta"`
}
