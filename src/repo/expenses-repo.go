package repo

import (
	"context"
	"interview-teamex-v1/src/db/model"
	"interview-teamex-v1/src/repo/opts"
)

func (repo *Repo) ExpenseFind(ctx context.Context, id int) (*model.Expense, error) {

	//
	var e model.Expense

	//
	err := repo.Conn.QueryRow(
		ctx, `SELECT id, date, amount, type, category FROM expenses WHERE id = $1`, id,
	).Scan(&e.ID, &e.Date, &e.Amount, &e.Type, &e.Category)
	if err != nil {
		return nil, err
	}

	//
	return &e, nil

}

func (repo *Repo) ExpenseCreate(
	ctx context.Context, date string, amount float64, tipo model.ExpenseType, category string,
) (*model.Expense, error) {
	//
	var id int
	err := repo.Conn.QueryRow(
		ctx,
		`INSERT INTO expenses (date, amount, type, category) VALUES ($1, $2, $3, $4) RETURNING id`,
		date, amount, tipo, category,
	).Scan(&id)

	// something wrong?
	if err != nil {
		return nil, err
	}

	//
	return repo.ExpenseFind(ctx, id)
}

func (repo *Repo) ExpenseUpdate(
	ctx context.Context, date string, amount float64, tipo model.ExpenseType, category string, id int,
) error {
	sql := `
		UPDATE expenses 
        SET date = $1, amount = $2, type = $3, category = $4 
        WHERE id = $5
	`
	_, err := repo.Conn.Exec(ctx, sql, date, amount, tipo, category, id)
	return err
}

func (repo *Repo) ExpenseGetSummary(ctx context.Context, options []opts.ExpenseGetFilterOpt) (float64, float64, error) {

	//
	options = append(options, opts.WithSummary())
	var sql, whereParams = buildGetSql(options)

	//
	var totalIncome float64
	var totalExpense float64
	err := repo.Conn.QueryRow(ctx, sql, whereParams...).Scan(&totalIncome, &totalExpense)
	if err != nil {
		return 0, 0, err
	}

	//
	return totalIncome, totalExpense, nil

}

func (repo *Repo) ExpenseGet(ctx context.Context, options []opts.ExpenseGetFilterOpt) ([]model.Expense, error) {

	//
	var sql, whereParams = buildGetSql(options)

	// It is the list, not the summary... lets go
	rows, err := repo.Conn.Query(ctx, sql, whereParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.Expense, 0)
	for rows.Next() {
		var e model.Expense
		err := rows.Scan(&e.ID, &e.Date, &e.Amount, &e.Type, &e.Category)
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	// some checking
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	//
	return result, nil
}
