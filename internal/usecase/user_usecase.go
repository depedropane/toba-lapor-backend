package usecase

import (
	"errors"
	"toba-lapor-backend/internal/model"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/repository"
	"toba-lapor-backend/pkg/utils"
)

type UserUsecase interface {
	CreateAdminDinas(req dto.CreateAdminDinasRequest) (*dto.AdminDinasResponse, error)
	GetAllAdminDinas() ([]dto.AdminDinasResponse, error)
	UpdateAdminDinas(id uint, req dto.UpdateAdminDinasRequest) (*dto.AdminDinasResponse, error)
	ToggleUserStatus(id uint, isActive bool) error
	GetAllUsers() ([]dto.MasyarakatResponse, error)
}

type userUsecase struct {
	userRepo   repository.UserRepository
	roleRepo   repository.RoleRepository
	agencyRepo repository.AgencyRepository
}

func NewUserUsecase(userRepo repository.UserRepository, roleRepo repository.RoleRepository, agencyRepo repository.AgencyRepository) UserUsecase {
	return &userUsecase{userRepo, roleRepo, agencyRepo}
}

func (u *userUsecase) CreateAdminDinas(req dto.CreateAdminDinasRequest) (*dto.AdminDinasResponse, error) {
	existingUser, _ := u.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	agency, err := u.agencyRepo.FindByID(req.AgencyID)
	if err != nil {
		return nil, errors.New("agency not found")
	}

	role, err := u.roleRepo.FindByName("admin_dinas")
	if err != nil {
		return nil, errors.New("role admin_dinas not found in database")
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
		AgencyID: &req.AgencyID,
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return &dto.AdminDinasResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		IsActive:   user.IsActive,
		AgencyID:   *user.AgencyID,
		AgencyName: agency.Name,
	}, nil
}

func (u *userUsecase) GetAllAdminDinas() ([]dto.AdminDinasResponse, error) {
	users, err := u.userRepo.FindAllAdminDinas()
	if err != nil {
		return nil, err
	}

	var res []dto.AdminDinasResponse
	for _, user := range users {
		agencyName := ""
		if user.Agency != nil {
			agencyName = user.Agency.Name
		}
		
		res = append(res, dto.AdminDinasResponse{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			Phone:      user.Phone,
			IsActive:   user.IsActive,
			AgencyID:   *user.AgencyID,
			AgencyName: agencyName,
		})
	}
	return res, nil
}

func (u *userUsecase) UpdateAdminDinas(id uint, req dto.UpdateAdminDinasRequest) (*dto.AdminDinasResponse, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Role.Name != "admin_dinas" {
		return nil, errors.New("user is not an admin dinas")
	}

	agency, err := u.agencyRepo.FindByID(req.AgencyID)
	if err != nil {
		return nil, errors.New("agency not found")
	}

	user.Name = req.Name
	user.Phone = req.Phone
	user.AgencyID = &req.AgencyID

	err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return &dto.AdminDinasResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		IsActive:   user.IsActive,
		AgencyID:   *user.AgencyID,
		AgencyName: agency.Name,
	}, nil
}

func (u *userUsecase) ToggleUserStatus(id uint, isActive bool) error {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	// Jangan izinkan blokir super_admin
	if user.Role.Name == "super_admin" {
		return errors.New("cannot change status of super_admin")
	}

	user.IsActive = isActive
	return u.userRepo.Update(user)
}

func (u *userUsecase) GetAllUsers() ([]dto.MasyarakatResponse, error) {
	users, err := u.userRepo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	var res []dto.MasyarakatResponse
	for _, user := range users {
		res = append(res, dto.MasyarakatResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Phone:    user.Phone,
			IsActive: user.IsActive,
		})
	}
	return res, nil
}
