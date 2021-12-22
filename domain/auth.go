package domain

import "github.com/mstreet3/banking-auth/errs"

type AuthRepository interface {
	Login(Login) (*Login, *errs.AppError)
}
