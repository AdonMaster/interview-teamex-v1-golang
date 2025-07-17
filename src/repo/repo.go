package repo

import (
	"github.com/jackc/pgx/v5"
)

type Repo struct {
	Conn *pgx.Conn
}

func New(conn *pgx.Conn) Repo {
	return Repo{conn}
}
