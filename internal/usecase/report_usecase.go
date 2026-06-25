package usecase

import (
	"errors"
	"toba-lapor-backend/internal/model"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/repository"
	"toba-lapor-backend/internal/service"
)

type ReportUsecase interface {
	CreateReport(req dto.CreateReportRequest, userID uint, imageUrls []string) (*dto.ReportResponse, error)
	GetMyReports(userID uint) ([]dto.ReportResponse, error)
	GetReportByID(id uint, userID uint, role string) (*dto.ReportResponse, error)
	GetAllReports() ([]dto.ReportResponse, error)
	VerifyReport(id uint, req dto.VerifyReportRequest, superAdminID uint) error
	RejectReport(id uint, req dto.RejectReportRequest, superAdminID uint) error
}

type reportUsecase struct {
	reportRepo  repository.ReportRepository
	historyRepo repository.ReportHistoryRepository
	agencyRepo  repository.AgencyRepository
	notifRepo   repository.NotificationRepository
	userRepo    repository.UserRepository
	firebaseSvc service.FirebaseService
}

func NewReportUsecase(reportRepo repository.ReportRepository, historyRepo repository.ReportHistoryRepository, agencyRepo repository.AgencyRepository, notifRepo repository.NotificationRepository, userRepo repository.UserRepository, firebaseSvc service.FirebaseService) ReportUsecase {
	return &reportUsecase{reportRepo, historyRepo, agencyRepo, notifRepo, userRepo, firebaseSvc}
}

func (u *reportUsecase) CreateReport(req dto.CreateReportRequest, userID uint, imageUrls []string) (*dto.ReportResponse, error) {
	var images []model.ReportImage
	for _, url := range imageUrls {
		images = append(images, model.ReportImage{
			ImageURL: url,
		})
	}

	report := &model.Report{
		Title:        req.Title,
		Description:  req.Description,
		Location:     req.Location,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		Status:       "PENDING_VERIFICATION",
		UserID:       userID,
		AgencyID:     req.AgencyID,
		ReportImages: images,
	}

	err := u.reportRepo.Create(report)
	if err != nil {
		return nil, err
	}

	// Create Initial History
	history := &model.ReportHistory{
		ReportID: report.ID,
		UserID:   userID,
		Status:   "PENDING_VERIFICATION",
		Notes:    "Laporan baru dibuat dan menunggu verifikasi",
	}
	u.historyRepo.Create(history)

	return u.GetReportByID(report.ID, userID, "user") // Reuse getter to construct response
}

func (u *reportUsecase) GetMyReports(userID uint) ([]dto.ReportResponse, error) {
	reports, err := u.reportRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var res []dto.ReportResponse
	for _, r := range reports {
		res = append(res, *mapToReportResponse(&r))
	}
	return res, nil
}

func (u *reportUsecase) GetReportByID(id uint, userID uint, role string) (*dto.ReportResponse, error) {
	report, err := u.reportRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("report not found")
	}

	// Authorization: User can only see their own reports.
	if role == "user" && report.UserID != userID {
		return nil, errors.New("unauthorized to view this report")
	}

	return mapToReportResponse(report), nil
}

func (u *reportUsecase) GetAllReports() ([]dto.ReportResponse, error) {
	reports, err := u.reportRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.ReportResponse
	for _, r := range reports {
		res = append(res, *mapToReportResponse(&r))
	}
	return res, nil
}

func (u *reportUsecase) VerifyReport(id uint, req dto.VerifyReportRequest, superAdminID uint) error {
	report, err := u.reportRepo.FindByID(id)
	if err != nil {
		return errors.New("report not found")
	}

	if report.Status != "PENDING_VERIFICATION" {
		return errors.New("report is not pending verification")
	}

	agency, err := u.agencyRepo.FindByID(req.AgencyID)
	if err != nil {
		return errors.New("agency not found")
	}

	report.AgencyID = &agency.ID
	report.Status = "ASSIGNED"

	err = u.reportRepo.Update(report)
	if err != nil {
		return err
	}

	history := &model.ReportHistory{
		ReportID: report.ID,
		UserID:   superAdminID,
		Status:   "ASSIGNED",
		Notes:    "Laporan diverifikasi dan diteruskan ke " + agency.Name,
	}
	u.historyRepo.Create(history)

	notif := &model.Notification{
		UserID: report.UserID,
		Title:  "Laporan Diverifikasi",
		Body:   "Laporan Anda tentang '" + report.Title + "' telah diteruskan ke " + agency.Name,
	}
	u.notifRepo.Create(notif)

	// Send Push Notification
	user, err := u.userRepo.FindByID(report.UserID)
	if err == nil && user.FCMToken != "" {
		u.firebaseSvc.SendPushNotification(user.FCMToken, notif.Title, notif.Body)
	}

	return nil
}

func (u *reportUsecase) RejectReport(id uint, req dto.RejectReportRequest, superAdminID uint) error {
	report, err := u.reportRepo.FindByID(id)
	if err != nil {
		return errors.New("report not found")
	}

	if report.Status != "PENDING_VERIFICATION" {
		return errors.New("report is not pending verification")
	}

	report.Status = "REJECTED"

	err = u.reportRepo.Update(report)
	if err != nil {
		return err
	}

	history := &model.ReportHistory{
		ReportID: report.ID,
		UserID:   superAdminID,
		Status:   "REJECTED",
		Notes:    req.Reason,
	}
	u.historyRepo.Create(history)

	notif := &model.Notification{
		UserID: report.UserID,
		Title:  "Laporan Ditolak",
		Body:   "Laporan Anda tentang '" + report.Title + "' ditolak. Alasan: " + req.Reason,
	}
	u.notifRepo.Create(notif)

	// Send Push Notification
	user, err := u.userRepo.FindByID(report.UserID)
	if err == nil && user.FCMToken != "" {
		u.firebaseSvc.SendPushNotification(user.FCMToken, notif.Title, notif.Body)
	}

	return nil
}

func mapToReportResponse(r *model.Report) *dto.ReportResponse {
	var images []dto.ReportImageResponse
	for _, img := range r.ReportImages {
		images = append(images, dto.ReportImageResponse{
			ID:       img.ID,
			ImageURL: img.ImageURL,
		})
	}

	var histories []dto.ReportHistoryResponse
	for _, h := range r.ReportHistories {
		histories = append(histories, dto.ReportHistoryResponse{
			ID:        h.ID,
			Status:    h.Status,
			Notes:     h.Notes,
			CreatedAt: h.CreatedAt,
		})
	}

	agencyName := ""
	if r.Agency != nil {
		agencyName = r.Agency.Name
	}

	return &dto.ReportResponse{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
		Location:    r.Location,
		Latitude:    r.Latitude,
		Longitude:   r.Longitude,
		Status:      r.Status,
		UserID:      r.UserID,
		UserName:    r.User.Name,
		AgencyID:    r.AgencyID,
		AgencyName:  agencyName,
		Images:      images,
		Histories:   histories,
		CreatedAt:   r.CreatedAt,
	}
}
