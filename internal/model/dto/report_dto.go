package dto

import "time"

type CreateReportRequest struct {
	Title       string  `form:"title" binding:"required"`
	Description string  `form:"description" binding:"required"`
	Location    string  `form:"location" binding:"required"`
	Latitude    float64 `form:"latitude"`
	Longitude   float64 `form:"longitude"`
	AgencyID    *uint   `form:"agency_id"` 
}

type ReportResponse struct {
	ID          uint                 `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Location    string               `json:"location"`
	Latitude    float64              `json:"latitude"`
	Longitude   float64              `json:"longitude"`
	Status      string               `json:"status"`
	UserID      uint                 `json:"user_id"`
	UserName    string               `json:"user_name"`
	AgencyID    *uint                `json:"agency_id"`
	AgencyName  string                  `json:"agency_name"`
	Images      []ReportImageResponse   `json:"images"`
	Histories   []ReportHistoryResponse `json:"histories"`
	CreatedAt   time.Time               `json:"created_at"`
}

type ReportImageResponse struct {
	ID       uint   `json:"id"`
	ImageURL string `json:"image_url"`
}

type ReportHistoryResponse struct {
	ID        uint      `json:"id"`
	Status    string    `json:"status"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}

type VerifyReportRequest struct {
	AgencyID uint `json:"agency_id" binding:"required"`
}

type RejectReportRequest struct {
	Reason string `json:"reason" binding:"required"`
}

type DashboardResponse struct {
	TotalReports         int64           `json:"total_reports"`
	ActiveReports        int64           `json:"active_reports"`
	CompletedReports     int64           `json:"completed_reports"`
	DistributionByAgency []AgencyStatDto `json:"distribution_by_agency"`
}

type AgencyStatDto struct {
	AgencyID   uint   `json:"agency_id"`
	AgencyName string `json:"agency_name"`
	Count      int64  `json:"count"`
}
