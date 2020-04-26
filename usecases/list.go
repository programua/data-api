package usecases

import "github.com/jmoiron/sqlx"

type ListUseCases interface {
	List(db *sqlx.DB)
}

func List(lu ListUseCases, db *sqlx.DB) {
	lu.List(db)
}
