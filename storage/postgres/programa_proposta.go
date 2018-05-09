package postgres

import (
	"fmt"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/storage"
)

func (dao *postgres) CreateProgramaProposta(programaProposta *schema.ProgramaProposta) (string, error) {
	if _, err := dao.db.Model(programaProposta).Insert(); err != nil {
		return "", err
	}

	return programaProposta.ID, nil
}

func (dao *postgres) ListProgramaProposta(idPrograma string, pagination *storage.Pagination) ([]*schema.Proposta, int, error) {
	if idPrograma == "" {
		return nil, 0, fmt.Errorf("id_programa is required")
	}

	sql := `
		SELECT
			propostas.id,
			propostas.id_proposta,
			propostas.uf_proponente,
			propostas.munic_proponente,
			propostas.cod_munic_ibge,
			propostas.cod_orgao_sup,
			propostas.desc_orgao_sup,
			propostas.natureza_juridica,
			propostas.nr_proposta,
			propostas.dia_proposta,
			propostas.cod_orgao,
			propostas.desc_orgao,
			propostas.modalidade,
			propostas.identif_proponente,
			propostas.nm_proponente,
			propostas.cep_proponente,
			propostas.endereco_proponente,
			propostas.bairro_proponente,
			propostas.nm_banco,
			propostas.situacao_conta,
			propostas.situacao_projeto_basico,
			propostas.sit_proposta,
			propostas.dia_inic_vigencia_proposta,
			propostas.dia_fim_vigencia_proposta,
			propostas.objeto_proposta,
			propostas.vl_global_prop,
			propostas.vl_repasse_prop,
			propostas.vl_contrapartida_prop,
			propostas.created_at,
			propostas.updated_at
		FROM propostas
			INNER JOIN programa_propostas ON programa_propostas.id_proposta = propostas.id_proposta
		WHERE
		  	programa_propostas.id_programa = ?
		ORDER BY propostas.uf_proponente, propostas.id_proposta, propostas.munic_proponente
	`

	countSql := "SELECT count(*) FROM programa_propostas WHERE id_programa = ?"
	models := []*schema.Proposta{nil}
	_, total, err := dao.query(&models, sql, countSql, pagination, idPrograma)
	return models, total, err
}
