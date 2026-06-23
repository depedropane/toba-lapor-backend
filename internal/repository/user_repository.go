package repository

import (
	"toba-lapor-backend/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	FindAllAdminDinas() ([]model.User, error)
	FindAllUsers() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Role").Preload("Agency").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) FindAllAdminDinas() ([]model.User, error) {
	var users []model.User
	err := r.db.Preload("Role").Preload("Agency").
		Joins("JOIN roles ON roles.id = users.role_id").
		Where("roles.name = ?", "admin_dinas").
		Find(&users).Error
	return users, err
}

func (r *userRepository) FindAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.db.Preload("Role").
		Joins("JOIN roles ON roles.id = users.role_id").
		Where("roles.name = ?", "user").
		Find(&users).Error
	return users, err
}
