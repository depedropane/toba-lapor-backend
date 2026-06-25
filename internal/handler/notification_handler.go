package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"toba-lapor-backend/internal/usecase"
	"toba-lapor-backend/pkg/utils"
)

type NotificationHandler struct {
	notifUsecase usecase.NotificationUsecase
}

func NewNotificationHandler(notifUsecase usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{notifUsecase}
}

func (h *NotificationHandler) GetMyNotifications(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	userID := uint(userIDRaw.(float64))

	res, err := h.notifUsecase.GetMyNotifications(userID)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	userID := uint(userIDRaw.(float64))

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", "Invalid ID")
		return
	}

	err = h.notifUsecase.MarkAsRead(uint(id), userID)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Failed to mark as read", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", nil)
}
