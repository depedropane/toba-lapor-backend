package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"toba-lapor-backend/internal/usecase"
	"toba-lapor-backend/pkg/utils"
)

type DashboardHandler struct {
	dashboardUsecase usecase.DashboardUsecase
}

func NewDashboardHandler(dashboardUsecase usecase.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{dashboardUsecase}
}

func (h *DashboardHandler) GetStats(c *gin.Context) {
	res, err := h.dashboardUsecase.GetDashboardStats()
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}
