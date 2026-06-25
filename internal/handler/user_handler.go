package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/usecase"
	"toba-lapor-backend/pkg/utils"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) CreateAdminDinas(c *gin.Context) {
	var req dto.CreateAdminDinasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	res, err := h.userUsecase.CreateAdminDinas(req)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Failed to create Admin Dinas", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusCreated, "Success", res)
}

func (h *UserHandler) GetAllAdminDinas(c *gin.Context) {
	res, err := h.userUsecase.GetAllAdminDinas()
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *UserHandler) UpdateAdminDinas(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", "Invalid ID")
		return
	}

	var req dto.UpdateAdminDinasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	res, err := h.userUsecase.UpdateAdminDinas(uint(id), req)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *UserHandler) ToggleUserStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", "Invalid ID")
		return
	}

	var req dto.UpdateUserStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	err = h.userUsecase.ToggleUserStatus(uint(id), *req.IsActive)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", nil)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	res, err := h.userUsecase.GetAllUsers()
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	userID := uint(userIDRaw.(float64))

	res, err := h.userUsecase.GetProfile(userID)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusNotFound, "Not Found", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	userID := uint(userIDRaw.(float64))

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	res, err := h.userUsecase.UpdateProfile(userID, req)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *UserHandler) UpdateFCMToken(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	userID := uint(userIDRaw.(float64))

	var req dto.UpdateFCMTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	err := h.userUsecase.UpdateFCMToken(userID, req)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", nil)
}
