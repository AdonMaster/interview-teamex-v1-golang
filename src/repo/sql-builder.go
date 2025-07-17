package repo

import (
	"fmt"
	"interview-teamex-v1/src/repo/opts"
)

func buildGetSql(options []opts.ExpenseGetFilterOpt) (string, []interface{}) {
	//
	ff := opts.ExpenseGetFilterOptFrom(options)
	page := 1
	pageSize := 10
	if ff.Page != nil {
		page = *ff.Page
	}
	if ff.PageSize != nil {
		pageSize = *ff.PageSize
	}

	//
	var where = "WHERE (1=1)"
	var whereParams []interface{}
	if ff.Category != nil {
		where += fmt.Sprintf(" AND (category LIKE $%d)", len(whereParams)+1)
		whereParams = append(whereParams, *ff.Category)
	}
	if ff.DtIni != nil {
		where += fmt.Sprintf(" AND (date >= $%d)", len(whereParams)+1)
		whereParams = append(whereParams, *ff.DtIni)
	}
	if ff.DtEnd != nil {
		where += fmt.Sprintf(" AND (date <= $%d)", len(whereParams)+1)
		whereParams = append(whereParams, *ff.DtEnd)
	}

	// finally building the sql.... keep in mind if I would to manipulate this properly...
	// I would use param array + string interpolation...
	sql := fmt.Sprintf(
		"SELECT * FROM expenses %s order by date desc LIMIT %d OFFSET %d", where, pageSize, (page-1)*pageSize,
	)

	// where id (sortable column/field)

	// one last thing...
	// is this summary or not
	if ff.IsSummary {
		subQuerySql := `
			SELECT 
				COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0.0) AS total_income,
				COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0.0) AS total_expense
			FROM
				 (%s)
		`
		sql = fmt.Sprintf(subQuerySql, sql)
	}

	return sql, whereParams
}
