package dto

import "github.com/mstreet3/banking-auth/errs"

type RegisterUserRequest struct {
	Username   string `json:"username"`
	CustomerId string `json:"customer_id"`
	Password   string `json:"password"`
}

type RegisterUserResponse struct {
	Username   string `json:"username"`
	CustomerId string `json:"customer_id"`
}

func (req RegisterUserRequest) Validate() *errs.AppError {
	if req.CustomerId == "" || req.Username == "" || req.Password == "" {
		return errs.NewValidationError("username, customer_id and password are required")
	}
	return nil
}
