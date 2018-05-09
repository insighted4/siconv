CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists consorcios
(
  id                     uuid default uuid_generate_v4()        not null,
  id_proposta            varchar(255),
  cnpj_consorcio         varchar(255),
  nome_consorcio         varchar(255),
  codigo_cnae_primario   varchar(255),
  desc_cnae_primario     text,
  codigo_cnae_secundario varchar(255),
  desc_cnae_secundario   text,
  cnpj_participante      varchar(255),
  nome_participante      varchar(255),
  file_ref               varchar(255),
  line_ref               integer,
  created_at             timestamp with time zone default now() not null,
  updated_at             timestamp with time zone default now() not null,
  constraint consorcios_pkey
  primary key (id),
  constraint consorcios_id_proposta_cnpj_consorcio_nome_consorcio_codigo_cna
  unique (id_proposta, cnpj_consorcio, nome_consorcio, codigo_cnae_primario, codigo_cnae_secundario, cnpj_participante),
  constraint consorcios_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists consorcios_id_proposta_index
  on consorcios (id_proposta);

create index if not exists consorcios_cnpj_consorcio_index
  on consorcios (cnpj_consorcio);

create index if not exists consorcios_cnpj_participante_index
  on consorcios (cnpj_participante);

create table if not exists convenios
(
  id                        uuid default uuid_generate_v4()        not null,
  nr_convenio               varchar(255),
  id_proposta               varchar(255),
  dia_assin_conv            date,
  sit_convenio              varchar(255),
  subsituacao_conv          varchar(255),
  situacao_publicacao       varchar(255),
  instrumento_ativo         boolean,
  ind_opera_obtv            boolean,
  nr_processo               varchar(255),
  ug_emitente               varchar(255),
  dia_publ_conv             varchar(255),
  dia_inic_vigenc_conv      date,
  dia_fim_vigenc_conv       date,
  dias_prest_contas         date,
  dia_limite_prest_contas   date,
  situacao_contratacao      varchar(255),
  ind_assinado              boolean,
  qtde_convenios            varchar(255),
  qtd_ta                    varchar(255),
  qtd_prorroga              varchar(255),
  vl_global_conv            real,
  vl_repasse_conv           real,
  vl_contrapartida_conv     real,
  vl_empenhado_conv         real,
  vl_desembolsado_conv      real,
  vl_saldo_reman_tesouro    real,
  vl_saldo_reman_convenente real,
  vl_rendimento_aplicacao   real,
  vl_ingresso_contrapartida real,
  file_ref                  varchar(255),
  line_ref                  integer,
  created_at                timestamp with time zone default now() not null,
  updated_at                timestamp with time zone default now() not null,
  constraint convenios_pkey
  primary key (id),
  constraint convenios_nr_convenio_unique
  unique (nr_convenio),
  constraint convenios_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists convenios_id_proposta_index
  on convenios (id_proposta);

create index if not exists convenios_dia_assin_conv_index
  on convenios (dia_assin_conv);

create index if not exists convenios_nr_processo_index
  on convenios (nr_processo);

create table if not exists desembolsos
(
  id                      uuid default uuid_generate_v4()        not null,
  nr_convenio             varchar(255),
  dt_ult_desembolso       date,
  qtd_dias_sem_desembolso integer,
  id_desembolso           varchar(255),
  data_desembolso         date,
  ano_desembolso          integer,
  mes_desembolso          integer,
  nr_siafi                varchar(255),
  vl_desembolsado         real,
  file_ref                varchar(255),
  line_ref                integer,
  created_at              timestamp with time zone default now() not null,
  updated_at              timestamp with time zone default now() not null,
  constraint desembolsos_pkey
  primary key (id),
  constraint desembolsos_nr_convenio_id_desembolso_unique
  unique (nr_convenio, id_desembolso),
  constraint desembolsos_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists desembolsos_nr_convenio_index
  on desembolsos (nr_convenio);

create table if not exists emendas
(
  id                            uuid default uuid_generate_v4()        not null,
  id_proposta                   varchar(255),
  qualif_proponente             varchar(255),
  cod_programa_emenda           varchar(255),
  nr_emenda                     varchar(255),
  nome_parlamentar              varchar(255),
  beneficiario_emenda           varchar(255),
  ind_impositivo                boolean,
  tipo_parlamentar              varchar(255),
  valor_repasse_proposta_emenda real,
  valor_repasse_emenda          real,
  file_ref                      varchar(255),
  line_ref                      integer,
  created_at                    timestamp with time zone default now() not null,
  updated_at                    timestamp with time zone default now() not null,
  constraint emendas_pkey
  primary key (id),
  constraint emendas_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists emendas_id_proposta_index
  on emendas (id_proposta);

create index if not exists emendas_cod_programa_emenda_index
  on emendas (cod_programa_emenda);

create index if not exists emendas_nr_emenda_index
  on emendas (nr_emenda);

create table if not exists empenhos
(
  id                    uuid default uuid_generate_v4()        not null,
  id_empenho            varchar(255),
  nr_empenho            varchar(255),
  tipo_nota             varchar(255),
  desc_tipo_nota        text,
  data_emissao          varchar(255),
  cod_situacao_empenho  varchar(255),
  desc_situacao_empenho text,
  valor_empenho         varchar(255),
  nr_convenio           varchar(255),
  file_ref              varchar(255),
  line_ref              integer,
  created_at            timestamp with time zone default now() not null,
  updated_at            timestamp with time zone default now() not null,
  constraint empenhos_pkey
  primary key (id),
  constraint empenhos_id_empenho_unique
  unique (id_empenho),
  constraint empenhos_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists empenhos_nr_empenho_index
  on empenhos (nr_empenho);

create table if not exists empenho_desembolsos
(
  id            uuid default uuid_generate_v4()        not null,
  id_desembolso varchar(255),
  id_empenho    varchar(255),
  valor_grupo   real,
  file_ref      varchar(255),
  line_ref      integer,
  created_at    timestamp with time zone default now() not null,
  updated_at    timestamp with time zone default now() not null,
  constraint empenho_desembolsos_pkey
  primary key (id),
  constraint empenho_desembolsos_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists empenho_desembolsos_id_desembolso_id_empenho_index
  on empenho_desembolsos (id_desembolso, id_empenho);

create table if not exists etapa_crono_fisicos
(
  id                     uuid default uuid_generate_v4()        not null,
  id_meta                varchar(255),
  id_etapa               varchar(255),
  nr_etapa               varchar(255),
  desc_etapa             text,
  data_inicio_etapa      date,
  data_fim_etapa         date,
  uf_etapa               varchar(255),
  municipio_etapa        varchar(255),
  endereco_etapa         text,
  cep_etapa              varchar(255),
  qtd_etapa              integer,
  und_fornecimento_etapa varchar(255),
  vl_etapa               real,
  file_ref               varchar(255),
  line_ref               integer,
  created_at             timestamp with time zone default now() not null,
  updated_at             timestamp with time zone default now() not null,
  constraint etapa_crono_fisicos_pkey
  primary key (id),
  constraint etapa_crono_fisicos_id_etapa_unique
  unique (id_etapa),
  constraint etapa_crono_fisicos_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists etapa_crono_fisicos_id_meta_index
  on etapa_crono_fisicos (id_meta);

create index if not exists etapa_crono_fisicos_nr_etapa_index
  on etapa_crono_fisicos (nr_etapa);

create index if not exists etapa_crono_fisicos_uf_etapa_index
  on etapa_crono_fisicos (uf_etapa);

create table if not exists historico_situacoes
(
  id                 uuid default uuid_generate_v4()        not null,
  id_proposta        varchar(255),
  nr_convenio        varchar(255),
  dia_historico_sit  date,
  historico_sit      varchar(255),
  dias_historico_sit integer,
  cod_historico_sit  varchar(255),
  file_ref           varchar(255),
  line_ref           integer,
  created_at         timestamp with time zone default now() not null,
  updated_at         timestamp with time zone default now() not null,
  constraint historico_situacoes_pkey
  primary key (id),
  constraint historico_situacoes_id_proposta_dia_historico_sit_cod_historico
  unique (id_proposta, dia_historico_sit, cod_historico_sit),
  constraint historico_situacoes_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists historico_situacoes_id_proposta_index
  on historico_situacoes (id_proposta);

create index if not exists historico_situacoes_nr_convenio_index
  on historico_situacoes (nr_convenio);

create table if not exists ingresso_contrapartidas
(
  id                        uuid default uuid_generate_v4()        not null,
  nr_convenio               varchar(255),
  dt_ingresso_contrapartida date,
  vl_ingresso_contrapartida real,
  file_ref                  varchar(255),
  line_ref                  integer,
  created_at                timestamp with time zone default now() not null,
  updated_at                timestamp with time zone default now() not null,
  constraint ingresso_contrapartidas_pkey
  primary key (id),
  constraint ingresso_contrapartidas_nr_convenio_dt_ingresso_contrapartida_u
  unique (nr_convenio, dt_ingresso_contrapartida),
  constraint ingresso_contrapartidas_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists ingresso_contrapartidas_nr_convenio_index
  on ingresso_contrapartidas (nr_convenio);

create table if not exists meta_crono_fisicos
(
  id                    uuid default uuid_generate_v4()        not null,
  id_meta               varchar(255),
  nr_convenio           varchar(255),
  cod_programa          varchar(255),
  nome_programa         varchar(255),
  nr_meta               integer,
  tipo_meta             varchar(255),
  desc_meta             text,
  data_inicio_meta      date,
  data_fim_meta         date,
  uf_meta               varchar(255),
  municipio_meta        varchar(255),
  endereco_meta         text,
  cep_meta              varchar(255),
  qtd_meta              integer,
  und_fornecimento_meta varchar(255),
  vl_meta               real,
  file_ref              varchar(255),
  line_ref              integer,
  created_at            timestamp with time zone default now() not null,
  updated_at            timestamp with time zone default now() not null,
  constraint meta_crono_fisicos_pkey
  primary key (id),
  constraint meta_crono_fisicos_id_meta_unique
  unique (id_meta),
  constraint meta_crono_fisicos_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists meta_crono_fisicos_nr_convenio_index
  on meta_crono_fisicos (nr_convenio);

create index if not exists meta_crono_fisicos_cod_programa_index
  on meta_crono_fisicos (cod_programa);

create index if not exists meta_crono_fisicos_nome_programa_index
  on meta_crono_fisicos (nome_programa);

create index if not exists meta_crono_fisicos_nr_meta_index
  on meta_crono_fisicos (nr_meta);

create index if not exists meta_crono_fisicos_uf_meta_index
  on meta_crono_fisicos (uf_meta);

create table if not exists obtv_convenentes
(
  id                           uuid default uuid_generate_v4()        not null,
  nr_mov_fin                   varchar(255),
  identif_favorecido_obtv_conv varchar(255),
  nm_favorecido_obtv_conv      varchar(255),
  tp_aquisicao                 varchar(255),
  vl_pago_obtv_conv            real,
  file_ref                     varchar(255),
  line_ref                     integer,
  created_at                   timestamp with time zone default now() not null,
  updated_at                   timestamp with time zone default now() not null,
  constraint obtv_convenentes_pkey
  primary key (id),
  constraint obtv_convenentes_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists obtv_convenentes_nr_mov_fin_index
  on obtv_convenentes (nr_mov_fin);

create index if not exists obtv_convenentes_identif_favorecido_obtv_conv_index
  on obtv_convenentes (identif_favorecido_obtv_conv);

create table if not exists pagamentos
(
  id                 uuid default uuid_generate_v4()        not null,
  nr_mov_fin         varchar(255),
  nr_convenio        varchar(255),
  identif_fornecedor varchar(255),
  nome_fornecedor    varchar(255),
  tp_mov_financeira  varchar(255),
  data_pag           date,
  nr_dl              text,
  desc_dl            text,
  vl_pago            real,
  file_ref           varchar(255),
  line_ref           integer,
  created_at         timestamp with time zone default now() not null,
  updated_at         timestamp with time zone default now() not null,
  constraint pagamentos_pkey
  primary key (id),
  constraint pagamentos_nr_mov_fin_unique
  unique (nr_mov_fin),
  constraint pagamentos_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists pagamentos_nr_convenio_index
  on pagamentos (nr_convenio);

create index if not exists pagamentos_identif_fornecedor_index
  on pagamentos (identif_fornecedor);

create table if not exists plano_aplicacao_detalhados
(
  id                  uuid default uuid_generate_v4()        not null,
  id_proposta         varchar(255),
  sigla               varchar(255),
  municipio           varchar(255),
  natureza_aquisicao  integer,
  descricao_item      text,
  cep_item            varchar(255),
  endereco_item       text,
  tipo_despesa_item   varchar(255),
  natureza_despesa    varchar(255),
  sit_item            varchar(255),
  qtd_item            integer,
  valor_unitario_item real,
  valor_total_item    real,
  file_ref            varchar(255),
  line_ref            integer,
  created_at          timestamp with time zone default now() not null,
  updated_at          timestamp with time zone default now() not null,
  constraint plano_aplicacao_detalhados_pkey
  primary key (id),
  constraint plano_aplicacao_detalhados_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists plano_aplicacao_detalhados_id_proposta_index
  on plano_aplicacao_detalhados (id_proposta);

create index if not exists plano_aplicacao_detalhados_sigla_index
  on plano_aplicacao_detalhados (sigla);

create table if not exists programas
(
  id                         uuid default uuid_generate_v4()        not null,
  cod_orgao_sup_programa     varchar(255),
  desc_orgao_sup_programa    text,
  id_programa                varchar(255),
  cod_programa               varchar(255),
  nome_programa              varchar(255),
  sit_programa               varchar(255),
  data_disponibilizacao      date,
  ano_disponibilizacao       integer,
  dt_prog_ini_receb_prop     date,
  dt_prog_fim_receb_prop     date,
  dt_prog_ini_emenda_par     date,
  dt_prog_fim_emenda_par     date,
  dt_prog_ini_benef_esp      date,
  dt_prog_fim_benef_esp      date,
  modalidade_programa        varchar(255),
  natureza_juridica_programa varchar(255),
  uf_programa                varchar(255),
  acao_orcamentaria          varchar(255),
  file_ref                   varchar(255),
  line_ref                   integer,
  created_at                 timestamp with time zone default now() not null,
  updated_at                 timestamp with time zone default now() not null,
  constraint programas_pkey
  primary key (id),
  constraint programas_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists programas_cod_orgao_sup_programa_index
  on programas (cod_orgao_sup_programa);

create index if not exists programas_id_programa_index
  on programas (id_programa);

create index if not exists programas_cod_programa_index
  on programas (cod_programa);

create index if not exists programas_uf_programa_index
  on programas (uf_programa);

create table if not exists programa_propostas
(
  id          uuid default uuid_generate_v4()        not null,
  id_programa varchar(255),
  id_proposta varchar(255),
  file_ref    varchar(255),
  line_ref    integer,
  created_at  timestamp with time zone default now() not null,
  updated_at  timestamp with time zone default now() not null,
  constraint programa_propostas_pkey
  primary key (id),
  constraint programa_propostas_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists programa_propostas_id_programa_id_proposta_index
  on programa_propostas (id_programa, id_proposta);

create table if not exists proponentes
(
  id                   uuid default uuid_generate_v4()        not null,
  identif_proponente   varchar(255),
  nm_proponente        varchar(255),
  municipio_proponente varchar(255),
  uf_proponente        varchar(255),
  endereco_proponente  text,
  bairro_proponente    varchar(255),
  cep_proponente       varchar(255),
  email_proponente     varchar(255),
  telefone_proponente  varchar(255),
  fax_proponente       varchar(255),
  file_ref             varchar(255),
  line_ref             integer,
  created_at           timestamp with time zone default now() not null,
  updated_at           timestamp with time zone default now() not null,
  constraint proponentes_pkey
  primary key (id),
  constraint proponentes_identif_proponente_unique
  unique (identif_proponente),
  constraint proponentes_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists proponentes_nm_proponente_index
  on proponentes (nm_proponente);

create index if not exists proponentes_uf_proponente_index
  on proponentes (uf_proponente);

create table if not exists propostas
(
  id                         uuid default uuid_generate_v4()        not null,
  id_proposta                varchar(255),
  uf_proponente              varchar(255),
  munic_proponente           varchar(255),
  cod_munic_ibge             varchar(255),
  cod_orgao_sup              varchar(255),
  desc_orgao_sup             text,
  natureza_juridica          varchar(255),
  nr_proposta                varchar(255),
  dia_proposta               date,
  cod_orgao                  varchar(255),
  desc_orgao                 text,
  modalidade                 varchar(255),
  identif_proponente         varchar(255),
  nm_proponente              varchar(255),
  cep_proponente             varchar(255),
  endereco_proponente        text,
  bairro_proponente          varchar(255),
  nm_banco                   varchar(255),
  situacao_conta             varchar(255),
  situacao_projeto_basico    varchar(255),
  sit_proposta               varchar(255),
  dia_inic_vigencia_proposta date,
  dia_fim_vigencia_proposta  date,
  objeto_proposta            text,
  vl_global_prop             real,
  vl_repasse_prop            real,
  vl_contrapartida_prop      real,
  file_ref                   varchar(255),
  line_ref                   integer,
  created_at                 timestamp with time zone default now() not null,
  updated_at                 timestamp with time zone default now() not null,
  constraint propostas_pkey
  primary key (id),
  constraint propostas_id_proposta_unique
  unique (id_proposta),
  constraint propostas_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists propostas_uf_proponente_index
  on propostas (uf_proponente);

create index if not exists propostas_nr_proposta_index
  on propostas (nr_proposta);

create index if not exists propostas_cod_orgao_index
  on propostas (cod_orgao);

create index if not exists propostas_nm_proponente_index
  on propostas (nm_proponente);

create table if not exists prorroga_oficios
(
  id                     uuid default uuid_generate_v4()        not null,
  nr_convenio            varchar(255),
  nr_prorroga            varchar(255),
  dt_inicio_prorroga     date,
  dt_fim_prorroga        date,
  dias_prorroga          integer,
  dt_assinatura_prorroga date,
  sit_prorroga           varchar(255),
  file_ref               varchar(255),
  line_ref               integer,
  created_at             timestamp with time zone default now() not null,
  updated_at             timestamp with time zone default now() not null,
  constraint prorroga_oficios_pkey
  primary key (id),
  constraint prorroga_oficios_nr_convenio_nr_prorroga_dt_inicio_prorroga_dt_
  unique (nr_convenio, nr_prorroga, dt_inicio_prorroga, dt_fim_prorroga, dias_prorroga, dt_assinatura_prorroga, sit_prorroga),
  constraint prorroga_oficios_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists prorroga_oficios_nr_convenio_index
  on prorroga_oficios (nr_convenio);

create index if not exists prorroga_oficios_nr_prorroga_index
  on prorroga_oficios (nr_prorroga);

create table if not exists termo_aditivos
(
  id                  uuid default uuid_generate_v4()        not null,
  nr_convenio         varchar(255),
  numero_ta           varchar(255),
  tipo_ta             varchar(255),
  vl_global_ta        real,
  vl_repasse_ta       real,
  vl_contrapartida_ta real,
  dt_assinatura_ta    date,
  dt_inicio_ta        date,
  dt_fim_ta           date,
  justificativa_ta    text,
  file_ref            varchar(255),
  line_ref            integer,
  created_at          timestamp with time zone default now() not null,
  updated_at          timestamp with time zone default now() not null,
  constraint termo_aditivos_pkey
  primary key (id),
  constraint termo_aditivos_nr_convenio_numero_ta_tipo_ta_dt_inicio_ta_justi
  unique (nr_convenio, numero_ta, tipo_ta, dt_inicio_ta, justificativa_ta),
  constraint termo_aditivos_file_ref_line_ref_unique
  unique (file_ref, line_ref)
);

create index if not exists termo_aditivos_nr_convenio_index
  on termo_aditivos (nr_convenio);

create index if not exists termo_aditivos_numero_ta_index
  on termo_aditivos (numero_ta);

create table if not exists history
(
  id         serial                                 not null,
  file       varchar(255)                           not null,
  line       varchar(255)                           not null,
  persisted  varchar(255)                           not null,
  created_at timestamp with time zone default now() not null,
  updated_at timestamp with time zone default now() not null,
  constraint history_pkey
  primary key (id),
  constraint history_file_line_unique
  unique (file, line)
);
