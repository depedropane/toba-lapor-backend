package usecase

import (
	"errors"
	"toba-lapor-backend/internal/model"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/repository"
	"toba-lapor-backend/pkg/utils"
)

type AuthUsecase interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
	Register(req dto.RegisterRequest) (*dto.UserResponse, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewAuthUsecase(userRepo repository.UserRepository, roleRepo repository.RoleRepository) AuthUsecase {
	return &authUsecase{userRepo, roleRepo}
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

func (u *authUsecase) Register(req dto.RegisterRequest) (*dto.UserResponse, error) {
	existingUser, _ := u.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	role, err := u.roleRepo.FindByName("user")
	if err != nil {
		return nil, errors.New("role user not found in database")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
		IsActive: true,
		RoleID:   role.ID,
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		RoleName: role.Name,
	}, nil
}
