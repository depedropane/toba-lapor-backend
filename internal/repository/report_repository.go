package repository

import (
	"toba-lapor-backend/internal/model"
	"toba-lapor-backend/internal/model/dto"
	"gorm.io/gorm"
)

type ReportRepository interface {
	Create(report *model.Report) error
	Update(report *model.Report) error
	FindAll() ([]model.Report, error)
	FindByUserID(userID uint) ([]model.Report, error)
	FindByID(id uint) (*model.Report, error)
	
	// Dashboard Stats
	CountTotal() (int64, error)
	CountActive() (int64, error)
	CountCompleted() (int64, error)
	GetDistributionByAgency() ([]dto.AgencyStatDto, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db}
}

func (r *reportRepository) Create(report *model.Report) error {
	return r.db.Create(report).Error
}

func (r *reportRepository) Update(report *model.Report) error {
	return r.db.Save(report).Error
}

func (r *reportRepository) FindAll() ([]model.Report, error) {
	var reports []model.Report
	err := r.db.Preload("ReportImages").Preload("Agency").Preload("User").Order("created_at desc").Find(&reports).Error
	return reports, err
}

func (r *reportRepository) FindByUserID(userID uint) ([]model.Report, error) {
	var reports []model.Report
	err := r.db.Preload("ReportImages").Preload("Agency").Preload("User").Where("user_id = ?", userID).Order("created_at desc").Find(&reports).Error
	return reports, err
}

func (r *reportRepository) FindByID(id uint) (*model.Report, error) {
	var report model.Report
	err := r.db.Preload("ReportImages").Preload("Agency").Preload("User").Preload("ReportHistories").First(&report, id).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

// Dashboard Aggregations

func (r *reportRepository) CountTotal() (int64, error) {
	var count int64
	err := r.db.Model(&model.Report{}).Count(&count).Error
	return count, err
}

func (r *reportRepository) CountActive() (int64, error) {
	var count int64
	err := r.db.Model(&model.Report{}).Where("status NOT IN ?", []string{"Selesai", "Ditolak", "COMPLETED", "REJECTED"}).Count(&count).Error
	return count, err
}

func (r *reportRepository) CountCompleted() (int64, error) {
	var count int64
	err := r.db.Model(&model.Report{}).Where("status IN ?", []string{"Selesai", "COMPLETED"}).Count(&count).Error
	return count, err
}

func (r *reportRepository) GetDistributionByAgency() ([]dto.AgencyStatDto, error) {
	var stats []dto.AgencyStatDto
	err := r.db.Model(&model.Report{}).
		Select("agencies.id as agency_id, agencies.name as agency_name, COUNT(reports.id) as count").
		Joins("JOIN agencies ON agencies.id = reports.agency_id").
		Where("reports.agency_id IS NOT NULL").
		Group("agencies.id, agencies.name").
		Scan(&stats).Error
	return stats, err
}
