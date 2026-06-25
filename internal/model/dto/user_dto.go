package dto

type CreateAdminDinasRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone"`
	AgencyID uint   `json:"agency_id" binding:"required"`
}

type AdminDinasResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	IsActive   bool   `json:"is_active"`
	AgencyID   uint   `json:"agency_id"`
	AgencyName string `json:"agency_name"`
}

type UpdateAdminDinasRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone"`
	AgencyID uint   `json:"agency_id" binding:"required"`
}

type UpdateUserStatusRequest struct {
	IsActive *bool `json:"is_active" binding:"required"`
}

type MasyarakatResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

type UpdateProfileRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
}

type UpdateFCMTokenRequest struct {
	FCMToken string `json:"fcm_token" binding:"required"`
}
