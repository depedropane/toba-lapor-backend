package repository

import (
	"toba-lapor-backend/internal/model"
	"gorm.io/gorm"
)

type AgencyRepository interface {
	FindAll() ([]model.Agency, error)
	FindByID(id uint) (*model.Agency, error)
	Create(agency *model.Agency) error
	Update(agency *model.Agency) error
	Delete(id uint) error
}

type agencyRepository struct {
	db *gorm.DB
}

func NewAgencyRepository(db *gorm.DB) AgencyRepository {
	return &agencyRepository{db}
}

func (r *agencyRepository) FindAll() ([]model.Agency, error) {
	var agencies []model.Agency
	err := r.db.Find(&agencies).Error
	return agencies, err
}

func (r *agencyRepository) FindByID(id uint) (*model.Agency, error) {
	var agency model.Agency
	err := r.db.First(&agency, id).Error
	if err != nil {
		return nil, err
	}
	return &agency, nil
}

func (r *agencyRepository) Create(agency *model.Agency) error {
	return r.db.Create(agency).Error
}

func (r *agencyRepository) Update(agency *model.Agency) error {
	return r.db.Save(agency).Error
}

func (r *agencyRepository) Delete(id uint) error {
	return r.db.Delete(&model.Agency{}, id).Error
}
