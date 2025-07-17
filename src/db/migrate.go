package db

import (
	context "context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func tableExists(ctx context.Context, conn *pgx.Conn, name string) (bool, error) {
	//
	query := `
		SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_schema = 'public'
              AND table_name = $1
        );`

	//
	var exists bool
	err := conn.QueryRow(ctx, query, name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("tableExists: %w", err)
	}

	return exists, nil
}

func migrate(ctx context.Context, conn *pgx.Conn) error {

	exists, err := tableExists(ctx, conn, "expenses")
	if err != nil {
		return err
	}
	if exists {
		return nil
	} // table already exists, I'll halt the migration process...

	//
	sql := `
		CREATE TYPE expense_type AS ENUM ('income', 'expense');
		
		CREATE TABLE expenses (
			id SERIAL PRIMARY KEY,
			date TIMESTAMP NOT NULL,
			amount NUMERIC(10, 2) NOT NULL,
			type expense_type NOT NULL,
			category TEXT NOT NULL
		);
		
		CREATE INDEX idx_expenses_date ON expenses (date);
		CREATE INDEX idx_expenses_type ON expenses (type);
		CREATE INDEX idx_expenses_category ON expenses (category);
	`
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	//
	return nil
}
