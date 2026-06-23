package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/usecase"
	"toba-lapor-backend/pkg/utils"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	res, err := h.authUsecase.Login(req)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}
