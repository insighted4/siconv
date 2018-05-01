package postgres

import (
	"fmt"

	"github.com/insighted4/siconv/schema"
	"github.com/insighted4/siconv/siconv"
)

func (dao *postgres) CreatePrograma(programa *schema.Programa) (string, error) {
	if _, err := dao.db.Model(programa).Insert(); err != nil {
		return "", err
	}

	return programa.ID, nil
}

func (dao *postgres) GetPrograma(id string) (*schema.Programa, error) {
	var model schema.Programa
	_, err := dao.get(&model, id)

	return &model, err
}

func (dao *postgres) ListPrograma(idPrograma string, pagination *siconv.Pagination) ([]*schema.Programa, int, error) {
	var args []interface{}
	where := "true"

	if idPrograma != "" {
		args = append(args, idPrograma)
		where = where + " AND id_programa = ?"
	}

	sql := fmt.Sprintf(`
		SELECT
			id,
			cod_orgao_sup_programa,
			desc_orgao_sup_programa,
			id_programa,
			cod_programa,
			nome_programa,
			sit_programa,
			data_disponibilizacao,
			ano_disponibilizacao,
			dt_prog_ini_receb_prop,
			dt_prog_fim_receb_prop,
			dt_prog_ini_emenda_par,
			dt_prog_fim_emenda_par,
			dt_prog_ini_benef_esp,
			dt_prog_fim_benef_esp,
			modalidade_programa,
			natureza_juridica_programa,
			uf_programa,
			acao_orcamentaria,
			created_at,
			updated_at
		FROM
			programas
		WHERE %s
		ORDER BY id_programa, data_disponibilizacao, uf_programa
	`, where)

	countSql := fmt.Sprintf("SELECT COUNT(*) FROM programas WHERE %s", where)
	models := []*schema.Programa{}
	_, total, err := dao.list(&models, sql, countSql, pagination, args...)
	return models, total, err
}
