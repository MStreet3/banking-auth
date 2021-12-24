package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/mstreet3/banking-auth/domain"
	"github.com/mstreet3/banking-auth/dto"
	"github.com/mstreet3/banking-auth/errs"
	"github.com/mstreet3/banking-auth/transforms"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	Register(dto.RegisterUserRequest) (*dto.RegisterUserResponse, *errs.AppError)
	ParseClaims(string) (*dto.ClaimsResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) AuthService {
	return DefaultAuthService{repo}
}

func (s DefaultAuthService) Register(req dto.RegisterUserRequest) (*dto.RegisterUserResponse, *errs.AppError) {

	user := transforms.ToUser(req)

	/* attempt to get credentials for user from data source */
	u, err := s.repo.AddUser(user)
	if err != nil {
		return nil, err
	}

	/* attempt to transform credentials into a response */
	resp := transforms.ToDto(*u)

	/* return response */
	return &resp, nil
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	/* attempt to get credentials for user from data source */
	l, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	/* attempt to transform credentials into a response */
	var resp *dto.LoginResponse
	resp, err = l.ToDto()
	if err != nil {
		return nil, err
	}

	/* return response */
	return resp, nil
}

func (s DefaultAuthService) ParseClaims(ss string) (*dto.ClaimsResponse, *errs.AppError) {
	/* parse token */
	token, err := jwt.ParseWithClaims(ss,
		&dto.ClaimsResponse{},
		func(token *jwt.Token) (interface{}, error) {
			return domain.MySigningKey, nil
		})
	if err != nil {
		return nil, errs.NewAuthenticationError(err.Error())
	}

	/* cast token and return */
	if claims, ok := token.Claims.(*dto.ClaimsResponse); ok && token.Valid {
		return claims, nil
	}
	return nil, errs.InvalidAccessTokenError()

}
