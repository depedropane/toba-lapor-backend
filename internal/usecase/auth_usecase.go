package usecase

import (
	"errors"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/repository"
	"toba-lapor-backend/pkg/utils"
)

type AuthUsecase interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{userRepo}
}

func (u *authUsecase) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("akun Anda telah dinonaktifkan oleh administrator")
	}

	token, err := utils.GenerateToken(user.ID, user.Role.Name, user.AgencyID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			RoleName: user.Role.Name,
		},
	}, nil
}
