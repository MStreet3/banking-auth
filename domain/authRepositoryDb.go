package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/mstreet3/banking-auth/errs"
	"github.com/mstreet3/banking-auth/logger"
)

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func NewAuthRepositoryDb(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{dbClient}
}

func (db AuthRepositoryDb) FindBy(username, password string) (*Login, *errs.AppError) {
	var login Login
	sqlVerify := `SELECT username, u.customer_id, role, group_concat(a.account_id) as account_numbers FROM users u
                  LEFT JOIN accounts a ON a.customer_id = u.customer_id
                  WHERE username = ? and password = ?
                  GROUP BY a.customer_id`
	err := db.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewAuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.UnexpectedDatabaseError()
		}
	}
	return &login, nil
}

func (db AuthRepositoryDb) AddUser(u User) (*User, *errs.AppError) {
	sql := "INSERT INTO users (username, password, role, customer_id) VALUES (?,?,?,?)"
	result, err := db.client.Exec(sql, u.Username, u.Password, u.Role, u.CustomerId)
	if err != nil {
		logger.Error("error while creating new user " + err.Error())
		return nil, errs.UnexpectedDatabaseError()
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting id of new user " + err.Error())
		return nil, errs.UnexpectedDatabaseError()
	}
	u.CustomerId = strconv.FormatInt(id, 10)
	return &u, nil
}
