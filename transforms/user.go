package transforms

import (
	"github.com/mstreet3/banking-auth/domain"
	"github.com/mstreet3/banking-auth/dto"
)

func ToUser(req dto.RegisterUserRequest) domain.User {
	u := domain.User{
		Username:   req.Username,
		Password:   req.Password,
		Role:       string(domain.CLIENT),
		CustomerId: req.CustomerId,
	}
	return u
}

func ToDto(u domain.User) dto.RegisterUserResponse {
	return dto.RegisterUserResponse{
		Username:   u.Username,
		CustomerId: u.CustomerId,
	}
}
