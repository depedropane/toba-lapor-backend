package dto

type CreateAgencyRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateAgencyRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type AgencyResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
