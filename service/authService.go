package service

import (
	"github.com/mstreet3/banking-auth/domain"
	"github.com/mstreet3/banking-auth/dto"
	"github.com/mstreet3/banking-auth/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) AuthService {
	return DefaultAuthService{repo}
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
