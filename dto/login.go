package dto

import "github.com/mstreet3/banking-auth/errs"

type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

/* validate the login request */
func (req LoginRequest) Validate() *errs.AppError {
	if req.Username == "" || req.Password == "" {
		return errs.NewValidationError("username and password are required")
	}
	return nil
}
