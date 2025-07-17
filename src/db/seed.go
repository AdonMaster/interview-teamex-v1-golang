package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"interview-teamex-v1/src/db/model"
	"time"
)

func isTablePopulated(ctx context.Context, conn *pgx.Conn) (bool, error) {
	//
	var hasRows bool
	err := conn.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM expenses LIMIT 1)").Scan(&hasRows)
	if err != nil {
		return false, fmt.Errorf("migrate: %w", err)
	}
	//
	return hasRows, nil
}

func seed(ctx context.Context, conn *pgx.Conn) error {
	//
	populated, err := isTablePopulated(ctx, conn)
	if err != nil {
		return err
	}

	// is it empty?
	if populated {
		return nil
	}

	// data
	now := time.Now()
	records := []model.Expense{
		{0, now.Add(-time.Hour * 24 * 9), 1200.75, model.ExpenseTypeIncome, "Freelance Project"},
		{0, now.Add(-time.Hour * 24 * 8), 75.20, model.ExpenseTypeExpense, "Dinner out"},
		{0, now.Add(-time.Hour * 24 * 7), 30.00, model.ExpenseTypeExpense, "Coffee"},
		{0, now.Add(-time.Hour * 24 * 6), 500.00, model.ExpenseTypeIncome, "Investment Dividend"},
		{0, now.Add(-time.Hour * 24 * 5), 120.00, model.ExpenseTypeExpense, "Books"},
		{0, now.Add(-time.Hour * 24 * 4), 40.50, model.ExpenseTypeExpense, "Public Transport"},
		{0, now.Add(-time.Hour * 24 * 3), 400.5, model.ExpenseTypeIncome, "Salary"},
		{0, now.Add(-time.Hour * 24 * 2), 101.0, model.ExpenseTypeExpense, "Groceries"},
		{0, now.Add(-time.Hour * 24 * 1), 15.0, model.ExpenseTypeExpense, "Movies"},
		{0, now.Add(-time.Hour * 24 * 0), 25.50, model.ExpenseTypeExpense, "Snacks"},
	}
	for _, r := range records {
		//
		sql := `INSERT INTO expenses (date, amount, type, category) VALUES ($1, $2, $3, $4)`
		// Db exec
		_, err := conn.Exec(ctx, sql, r.Date, r.Amount, r.Type, r.Category)
		if err != nil {
			return fmt.Errorf("seed: %w", err)
		}
	}

	// phew, done!
	return nil
}
