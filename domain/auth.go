package domain

import "github.com/mstreet3/banking-auth/errs"

type AuthRepository interface {
	FindBy(username, password string) (*Login, *errs.AppError)
	AddUser(User) (*User, *errs.AppError)
}
