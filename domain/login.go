package domain

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/mstreet3/banking-auth/dto"
	"github.com/mstreet3/banking-auth/errs"
)

const TOKEN_EXPIRATION = time.Hour

var mySigningKey = []byte("SECRET_KEY") // todo: get key from environment variable

type AuthRole string

const (
	CLIENT AuthRole = "user"
	ADMIN  AuthRole = "admin"
)

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"account_numbers"`
	Role       AuthRole       `db:"role"`
}

func (l Login) validate() *errs.AppError {
	if l.CustomerId.Valid && l.Accounts.Valid && l.Role == CLIENT {
		return nil
	}
	if l.Role == ADMIN {
		return nil
	}
	return errs.NewAuthenticationError("invalid authentication attempt")

}

func (l Login) generateClaims() (*jwt.MapClaims, *errs.AppError) {
	invalid := l.validate()
	if invalid != nil {
		return nil, invalid
	}

	claims := jwt.MapClaims{
		"username": l.Username,
		"role":     l.Role,
		"exp":      time.Now().Add(TOKEN_EXPIRATION).Unix(),
	}

	switch l.Role {
	case CLIENT:
		claims["accounts"] = strings.Split(l.Accounts.String, ",")
		claims["customerId"] = l.CustomerId.String
	default:
		break
	}
	return &claims, nil
}

func (l Login) generateJwt() (*string, *errs.AppError) {
	claims, err := l.generateClaims()

	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, signErr := token.SignedString(mySigningKey)
	if signErr != nil {
		return nil, errs.NewAuthenticationError(fmt.Sprintf("could not create jwt: %s", signErr.Error()))
	}
	return &ss, nil

}

func (l Login) ToDto() (*dto.LoginResponse, *errs.AppError) {
	accessToken, err := l.generateJwt()
	if err != nil {
		return nil, err
	}
	resp := &dto.LoginResponse{
		AccessToken: *accessToken,
	}
	return resp, nil
}
