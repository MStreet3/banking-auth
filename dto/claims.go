package dto

import "github.com/golang-jwt/jwt"

type ClaimsResponse struct {
	Username   string   `json:"username"`
	Role       string   `json:"role"`
	Accounts   []string `json:"accounts,omitempty"`
	CustomerId string   `json:"customer_id,omitempty"`
	jwt.StandardClaims
}
