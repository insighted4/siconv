CREATE TABLE IF NOT EXISTS consorcios
(
  id integer not null
    CONSTRAINT consorcios_pkey
    primary key,
  id_proposta varchar(255),
  cnpj_consorcio varchar(255),
  nome_consorcio varchar(255),
  codigo_cnae_primario varchar(255),
  desc_cnae_primario text,
  codigo_cnae_secundario varchar(255),
  desc_cnae_secundario text,
  cnpj_participante varchar(255),
  nome_participante varchar(255),
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null,
  CONSTRAINT consorcios_id_proposta_cnpj_consorcio_nome_consorcio_codigo_cna
  UNIQUE (id_proposta, cnpj_consorcio, nome_consorcio, codigo_cnae_primario, codigo_cnae_secundario, cnpj_participante)
)
;

CREATE INDEX IF NOT EXISTS consorcios_id_proposta_index
  on consorcios (id_proposta)
;

CREATE INDEX IF NOT EXISTS consorcios_cnpj_consorcio_index
  on consorcios (cnpj_consorcio)
;

CREATE INDEX IF NOT EXISTS consorcios_cnpj_participante_index
  on consorcios (cnpj_participante)
;

CREATE TABLE IF NOT EXISTS convenios
(
  id integer not null
    CONSTRAINT convenios_pkey
    primary key,
  nr_convenio varchar(255)
    CONSTRAINT convenios_nr_convenio_UNIQUE
    UNIQUE,
  id_proposta varchar(255),
  dia_assin_conv date,
  sit_convenio varchar(255),
  subsituacao_conv varchar(255),
  situacao_publicacao varchar(255),
  instrumento_ativo boolean,
  ind_opera_obtv boolean,
  nr_processo varchar(255),
  ug_emitente varchar(255),
  dia_publ_conv varchar(255),
  dia_inic_vigenc_conv date,
  dia_fim_vigenc_conv date,
  dias_prest_contas date,
  dia_limite_prest_contas date,
  situacao_contratacao varchar(255),
  ind_assinado boolean,
  qtde_convenios varchar(255),
  qtd_ta varchar(255),
  qtd_prorroga varchar(255),
  vl_global_conv real,
  vl_repasse_conv real,
  vl_contrapartida_conv real,
  vl_empenhado_conv real,
  vl_desembolsado_conv real,
  vl_saldo_reman_tesouro real,
  vl_saldo_reman_convenente real,
  vl_rendimento_aplicacao real,
  vl_ingresso_contrapartida real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS convenios_id_proposta_index
  on convenios (id_proposta)
;

CREATE INDEX IF NOT EXISTS convenios_dia_assin_conv_index
  on convenios (dia_assin_conv)
;

CREATE INDEX IF NOT EXISTS convenios_nr_processo_index
  on convenios (nr_processo)
;

CREATE TABLE IF NOT EXISTS desembolsos
(
  id integer not null
    CONSTRAINT desembolsos_pkey
    primary key,
  nr_convenio varchar(255),
  dt_ult_desembolso date,
  qtd_dias_sem_desembolso integer,
  id_desembolso varchar(255),
  data_desembolso date,
  ano_desembolso integer,
  mes_desembolso integer,
  nr_siafi varchar(255),
  vl_desembolsado real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null,
  CONSTRAINT desembolsos_nr_convenio_id_desembolso_UNIQUE
  UNIQUE (nr_convenio, id_desembolso)
)
;

CREATE INDEX IF NOT EXISTS desembolsos_nr_convenio_index
  on desembolsos (nr_convenio)
;

CREATE TABLE IF NOT EXISTS emendas
(
  id integer not null
    CONSTRAINT emendas_pkey
    primary key,
  id_proposta varchar(255),
  qualif_proponente varchar(255),
  cod_programa_emenda varchar(255),
  nr_emenda varchar(255),
  nome_parlamentar varchar(255),
  beneficiario_emenda varchar(255),
  ind_impositivo boolean,
  tipo_parlamentar varchar(255),
  valor_repasse_proposta_emenda real,
  valor_repasse_emenda real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS emendas_id_proposta_index
  on emendas (id_proposta)
;

CREATE INDEX IF NOT EXISTS emendas_cod_programa_emenda_index
  on emendas (cod_programa_emenda)
;

CREATE INDEX IF NOT EXISTS emendas_nr_emenda_index
  on emendas (nr_emenda)
;

CREATE TABLE IF NOT EXISTS empenhos
(
  id integer not null
    CONSTRAINT empenhos_pkey
    primary key,
  id_empenho varchar(255)
    CONSTRAINT empenhos_id_empenho_UNIQUE
    UNIQUE,
  nr_empenho varchar(255),
  tipo_nota varchar(255),
  desc_tipo_nota text,
  data_emissao varchar(255),
  cod_situacao_empenho varchar(255),
  desc_situacao_empenho text,
  valor_empenho varchar(255),
  nr_convenio varchar(255),
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS empenhos_nr_empenho_index
  on empenhos (nr_empenho)
;

CREATE TABLE IF NOT EXISTS empenho_desembolsos
(
  id integer not null
    CONSTRAINT empenho_desembolsos_pkey
    primary key,
  id_desembolso varchar(255),
  id_empenho varchar(255),
  valor_grupo real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS empenho_desembolsos_id_desembolso_id_empenho_index
  on empenho_desembolsos (id_desembolso, id_empenho)
;

CREATE TABLE IF NOT EXISTS etapa_crono_fisicos
(
  id integer not null
    CONSTRAINT etapa_crono_fisicos_pkey
    primary key,
  id_meta varchar(255),
  id_etapa varchar(255)
    CONSTRAINT etapa_crono_fisicos_id_etapa_UNIQUE
    UNIQUE,
  nr_etapa varchar(255),
  desc_etapa text,
  data_inicio_etapa date,
  data_fim_etapa date,
  uf_etapa varchar(255),
  municipio_etapa varchar(255),
  endereco_etapa text,
  cep_etapa varchar(255),
  qtd_etapa integer,
  und_fornecimento_etapa varchar(255),
  vl_etapa real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS etapa_crono_fisicos_id_meta_index
  on etapa_crono_fisicos (id_meta)
;

CREATE INDEX IF NOT EXISTS etapa_crono_fisicos_nr_etapa_index
  on etapa_crono_fisicos (nr_etapa)
;

CREATE INDEX IF NOT EXISTS etapa_crono_fisicos_uf_etapa_index
  on etapa_crono_fisicos (uf_etapa)
;

CREATE TABLE IF NOT EXISTS historico_situacoes
(
  id integer not null
    CONSTRAINT historico_situacoes_pkey
    primary key,
  id_proposta varchar(255),
  nr_convenio varchar(255),
  dia_historico_sit date,
  historico_sit varchar(255),
  dias_historico_sit integer,
  cod_historico_sit varchar(255),
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null,
  CONSTRAINT historico_situacoes_id_proposta_dia_historico_sit_cod_historico
  UNIQUE (id_proposta, dia_historico_sit, cod_historico_sit)
)
;

CREATE INDEX IF NOT EXISTS historico_situacoes_id_proposta_index
  on historico_situacoes (id_proposta)
;

CREATE INDEX IF NOT EXISTS historico_situacoes_nr_convenio_index
  on historico_situacoes (nr_convenio)
;

CREATE TABLE IF NOT EXISTS ingresso_contrapartidas
(
  id integer not null
    CONSTRAINT ingresso_contrapartidas_pkey
    primary key,
  nr_convenio varchar(255),
  dt_ingresso_contrapartida date,
  vl_ingresso_contrapartida real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null,
  CONSTRAINT ingresso_contrapartidas_nr_convenio_dt_ingresso_contrapartida_u
  UNIQUE (nr_convenio, dt_ingresso_contrapartida)
)
;

CREATE INDEX IF NOT EXISTS ingresso_contrapartidas_nr_convenio_index
  on ingresso_contrapartidas (nr_convenio)
;

CREATE TABLE IF NOT EXISTS meta_crono_fisicos
(
  id integer not null
    CONSTRAINT meta_crono_fisicos_pkey
    primary key,
  id_meta varchar(255)
    CONSTRAINT meta_crono_fisicos_id_meta_UNIQUE
    UNIQUE,
  nr_convenio varchar(255),
  cod_programa varchar(255),
  nome_programa text,
  nr_meta integer,
  tipo_meta varchar(255),
  desc_meta text,
  data_inicio_meta date,
  data_fim_meta date,
  uf_meta varchar(255),
  municipio_meta varchar(255),
  endereco_meta text,
  cep_meta varchar(255),
  qtd_meta integer,
  und_fornecimento_meta varchar(255),
  vl_meta real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS meta_crono_fisicos_nr_convenio_index
  on meta_crono_fisicos (nr_convenio)
;

CREATE INDEX IF NOT EXISTS meta_crono_fisicos_cod_programa_index
  on meta_crono_fisicos (cod_programa)
;

CREATE INDEX IF NOT EXISTS meta_crono_fisicos_nr_meta_index
  on meta_crono_fisicos (nr_meta)
;

CREATE INDEX IF NOT EXISTS meta_crono_fisicos_uf_meta_index
  on meta_crono_fisicos (uf_meta)
;

CREATE TABLE IF NOT EXISTS obtv_convenentes
(
  id integer not null
    CONSTRAINT obtv_convenentes_pkey
    primary key,
  nr_mov_fin varchar(255),
  identif_favorecido_obtv_conv varchar(255),
  nm_favorecido_obtv_conv varchar(255),
  tp_aquisicao varchar(255),
  vl_pago_obtv_conv real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS obtv_convenentes_nr_mov_fin_index
  on obtv_convenentes (nr_mov_fin)
;

CREATE INDEX IF NOT EXISTS obtv_convenentes_identif_favorecido_obtv_conv_index
  on obtv_convenentes (identif_favorecido_obtv_conv)
;

CREATE TABLE IF NOT EXISTS pagamentos
(
  id integer not null
    CONSTRAINT pagamentos_pkey
    primary key,
  nr_mov_fin varchar(255)
    CONSTRAINT pagamentos_nr_mov_fin_UNIQUE
    UNIQUE,
  nr_convenio varchar(255),
  identif_fornecedor varchar(255),
  nome_fornecedor varchar(255),
  tp_mov_financeira varchar(255),
  data_pag date,
  nr_dl text,
  desc_dl text,
  vl_pago real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS pagamentos_nr_convenio_index
  on pagamentos (nr_convenio)
;

CREATE INDEX IF NOT EXISTS pagamentos_identif_fornecedor_index
  on pagamentos (identif_fornecedor)
;

CREATE TABLE IF NOT EXISTS plano_aplicacao_detalhados
(
  id integer not null
    CONSTRAINT plano_aplicacao_detalhados_pkey
    primary key,
  id_proposta varchar(255),
  sigla varchar(255),
  municipio varchar(255),
  natureza_aquisicao integer,
  descricao_item text,
  cep_item varchar(255),
  endereco_item text,
  tipo_despesa_item varchar(255),
  natureza_despesa varchar(255),
  sit_item varchar(255),
  qtd_item integer,
  valor_unitario_item real,
  valor_total_item real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS plano_aplicacao_detalhados_id_proposta_index
  on plano_aplicacao_detalhados (id_proposta)
;

CREATE INDEX IF NOT EXISTS plano_aplicacao_detalhados_sigla_index
  on plano_aplicacao_detalhados (sigla)
;

CREATE TABLE IF NOT EXISTS programas
(
  id integer not null
    CONSTRAINT programas_pkey
    primary key,
  cod_orgao_sup_programa varchar(255),
  desc_orgao_sup_programa text,
  id_programa varchar(255),
  cod_programa varchar(255),
  nome_programa varchar(255),
  sit_programa varchar(255),
  data_disponibilizacao date,
  ano_disponibilizacao integer,
  dt_prog_ini_receb_prop date,
  dt_prog_fim_receb_prop date,
  dt_prog_ini_emenda_par date,
  dt_prog_fim_emenda_par date,
  dt_prog_ini_benef_esp date,
  dt_prog_fim_benef_esp date,
  modalidade_programa varchar(255),
  natureza_juridica_programa varchar(255),
  uf_programa varchar(255),
  acao_orcamentaria varchar(255),
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS programas_cod_orgao_sup_programa_index
  on programas (cod_orgao_sup_programa)
;

CREATE INDEX IF NOT EXISTS programas_id_programa_index
  on programas (id_programa)
;

CREATE INDEX IF NOT EXISTS programas_cod_programa_index
  on programas (cod_programa)
;

CREATE INDEX IF NOT EXISTS programas_uf_programa_index
  on programas (uf_programa)
;

CREATE TABLE IF NOT EXISTS programa_propostas
(
  id integer not null
    CONSTRAINT programa_propostas_pkey
    primary key,
  id_programa varchar(255),
  id_proposta varchar(255),
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS programa_propostas_id_programa_id_proposta_index
  on programa_propostas (id_programa, id_proposta)
;

CREATE TABLE IF NOT EXISTS proponentes
(
  id integer not null
    CONSTRAINT proponentes_pkey
    primary key,
  identif_proponente varchar(255)
    CONSTRAINT proponentes_identif_proponente_UNIQUE
    UNIQUE,
  nm_proponente varchar(255),
  municipio_proponente varchar(255),
  uf_proponente varchar(255),
  endereco_proponente text,
  bairro_proponente varchar(255),
  cep_proponente varchar(255),
  email_proponente varchar(255),
  telefone_proponente varchar(255),
  fax_proponente varchar(255),
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS proponentes_nm_proponente_index
  on proponentes (nm_proponente)
;

CREATE INDEX IF NOT EXISTS proponentes_uf_proponente_index
  on proponentes (uf_proponente)
;

CREATE TABLE IF NOT EXISTS propostas
(
  id integer not null
    CONSTRAINT propostas_pkey
    primary key,
  id_proposta varchar(255)
    CONSTRAINT propostas_id_proposta_UNIQUE
    UNIQUE,
  uf_proponente varchar(255),
  munic_proponente varchar(255),
  cod_munic_ibge varchar(255),
  cod_orgao_sup varchar(255),
  desc_orgao_sup text,
  natureza_juridica varchar(255),
  nr_proposta varchar(255),
  dia_proposta date,
  cod_orgao varchar(255),
  desc_orgao text,
  modalidade varchar(255),
  identif_proponente varchar(255),
  nm_proponente varchar(255),
  cep_proponente varchar(255),
  endereco_proponente text,
  bairro_proponente varchar(255),
  nm_banco varchar(255),
  situacao_conta varchar(255),
  situacao_projeto_basico varchar(255),
  sit_proposta varchar(255),
  dia_inic_vigencia_proposta date,
  dia_fim_vigencia_proposta date,
  objeto_proposta text,
  vl_global_prop real,
  vl_repasse_prop real,
  vl_contrapartida_prop real,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS propostas_uf_proponente_index
  on propostas (uf_proponente)
;

CREATE INDEX IF NOT EXISTS propostas_nr_proposta_index
  on propostas (nr_proposta)
;

CREATE INDEX IF NOT EXISTS propostas_cod_orgao_index
  on propostas (cod_orgao)
;

CREATE INDEX IF NOT EXISTS propostas_nm_proponente_index
  on propostas (nm_proponente)
;

CREATE TABLE IF NOT EXISTS prorroga_oficios
(
  id integer not null
    CONSTRAINT prorroga_oficios_pkey
    primary key,
  nr_convenio varchar(255),
  nr_prorroga varchar(255),
  dt_inicio_prorroga date,
  dt_fim_prorroga date,
  dias_prorroga integer,
  dt_assinatura_prorroga date,
  sit_prorroga varchar(255),
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS prorroga_oficios_nr_convenio_index
  on prorroga_oficios (nr_convenio)
;

CREATE INDEX IF NOT EXISTS prorroga_oficios_nr_prorroga_index
  on prorroga_oficios (nr_prorroga)
;

CREATE TABLE IF NOT EXISTS termo_aditivos
(
  id integer not null
    CONSTRAINT termo_aditivos_pkey
    primary key,
  nr_convenio varchar(255),
  numero_ta varchar(255),
  tipo_ta varchar(255),
  vl_global_ta real,
  vl_repasse_ta real,
  vl_contrapartida_ta real,
  dt_assinatura_ta date,
  dt_inicio_ta date,
  dt_fim_ta date,
  justificativa_ta text,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null
)
;

CREATE INDEX IF NOT EXISTS termo_aditivos_nr_convenio_index
  on termo_aditivos (nr_convenio)
;

CREATE INDEX IF NOT EXISTS termo_aditivos_numero_ta_index
  on termo_aditivos (numero_ta)
;

