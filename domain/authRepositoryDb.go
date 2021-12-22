package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/mstreet3/banking-auth/errs"
)

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func NewAuthRepositoryDb(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{dbClient}
}

func (db AuthRepositoryDb) Login(l Login) (*Login, *errs.AppError) {
	return nil, nil
}
