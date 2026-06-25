package repository

import (
	"toba-lapor-backend/internal/model"
	"gorm.io/gorm"
)

type ReportHistoryRepository interface {
	Create(history *model.ReportHistory) error
	FindByReportID(reportID uint) ([]model.ReportHistory, error)
}

type reportHistoryRepository struct {
	db *gorm.DB
}

func NewReportHistoryRepository(db *gorm.DB) ReportHistoryRepository {
	return &reportHistoryRepository{db}
}

func (r *reportHistoryRepository) Create(history *model.ReportHistory) error {
	return r.db.Create(history).Error
}

func (r *reportHistoryRepository) FindByReportID(reportID uint) ([]model.ReportHistory, error) {
	var histories []model.ReportHistory
	err := r.db.Where("report_id = ?", reportID).Order("created_at desc").Find(&histories).Error
	return histories, err
}
