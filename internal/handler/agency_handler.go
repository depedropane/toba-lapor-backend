package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/usecase"
	"toba-lapor-backend/pkg/utils"
)

type AgencyHandler struct {
	agencyUsecase usecase.AgencyUsecase
}

func NewAgencyHandler(agencyUsecase usecase.AgencyUsecase) *AgencyHandler {
	return &AgencyHandler{agencyUsecase}
}

func (h *AgencyHandler) Create(c *gin.Context) {
	var req dto.CreateAgencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	res, err := h.agencyUsecase.CreateAgency(req)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusCreated, "Success", res)
}

func (h *AgencyHandler) GetAll(c *gin.Context) {
	res, err := h.agencyUsecase.GetAllAgencies()
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}
	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *AgencyHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", "Invalid ID")
		return
	}

	var req dto.UpdateAgencyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	res, err := h.agencyUsecase.UpdateAgency(uint(id), req)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *AgencyHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", "Invalid ID")
		return
	}

	err = h.agencyUsecase.DeleteAgency(uint(id))
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", nil)
}
