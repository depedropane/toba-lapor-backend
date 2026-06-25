package usecase

import (
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/repository"
)

type DashboardUsecase interface {
	GetDashboardStats() (*dto.DashboardResponse, error)
}

type dashboardUsecase struct {
	reportRepo repository.ReportRepository
}

func NewDashboardUsecase(reportRepo repository.ReportRepository) DashboardUsecase {
	return &dashboardUsecase{reportRepo}
}

func (u *dashboardUsecase) GetDashboardStats() (*dto.DashboardResponse, error) {
	total, err := u.reportRepo.CountTotal()
	if err != nil {
		return nil, err
	}

	active, err := u.reportRepo.CountActive()
	if err != nil {
		return nil, err
	}

	completed, err := u.reportRepo.CountCompleted()
	if err != nil {
		return nil, err
	}

	distribution, err := u.reportRepo.GetDistributionByAgency()
	if err != nil {
		return nil, err
	}

	return &dto.DashboardResponse{
		TotalReports:         total,
		ActiveReports:        active,
		CompletedReports:     completed,
		DistributionByAgency: distribution,
	}, nil
}
