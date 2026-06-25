package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/usecase"
	"toba-lapor-backend/pkg/utils"
)

type ReportHandler struct {
	reportUsecase usecase.ReportUsecase
}

func NewReportHandler(reportUsecase usecase.ReportUsecase) *ReportHandler {
	return &ReportHandler{reportUsecase}
}

func (h *ReportHandler) Create(c *gin.Context) {
	// Parse user ID from middleware
	userIDRaw, exists := c.Get("user_id")
	if !exists {
		utils.BuildErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "User ID not found in token")
		return
	}
	userID := uint(userIDRaw.(float64))

	var req dto.CreateReportRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	// Handle File Uploads
	form, _ := c.MultipartForm()
	var imageUrls []string

	if form != nil && form.File != nil {
		files := form.File["images"]
		
		// Create uploads directory if not exists
		err := os.MkdirAll("uploads", os.ModePerm)
		if err != nil {
			utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", "Failed to create upload directory")
			return
		}

		for _, file := range files {
			// Generate unique filename
			filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
			filePath := filepath.Join("uploads", filename)
			
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				utils.BuildErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", "Failed to save file")
				return
			}
			
			// For local storage, URL is just the path
			imageUrls = append(imageUrls, "/"+filepath.ToSlash(filePath))
		}
	}

	res, err := h.reportUsecase.CreateReport(req, userID, imageUrls)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Failed to create report", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusCreated, "Success", res)
}

func (h *ReportHandler) GetMyReports(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	userID := uint(userIDRaw.(float64))

	res, err := h.reportUsecase.GetMyReports(userID)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Failed to get reports", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *ReportHandler) GetDetail(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	userID := uint(userIDRaw.(float64))
	roleRaw, _ := c.Get("role")
	role := roleRaw.(string)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", "Invalid ID")
		return
	}

	res, err := h.reportUsecase.GetReportByID(uint(id), userID, role)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusNotFound, "Not Found", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *ReportHandler) GetAllReports(c *gin.Context) {
	res, err := h.reportUsecase.GetAllReports()
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Failed to get reports", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", res)
}

func (h *ReportHandler) VerifyReport(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	superAdminID := uint(userIDRaw.(float64))

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", "Invalid ID")
		return
	}

	var req dto.VerifyReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	err = h.reportUsecase.VerifyReport(uint(id), req, superAdminID)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Failed to verify report", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", nil)
}

func (h *ReportHandler) RejectReport(c *gin.Context) {
	userIDRaw, _ := c.Get("user_id")
	superAdminID := uint(userIDRaw.(float64))

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", "Invalid ID")
		return
	}

	var req dto.RejectReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BuildErrorResponse(c, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	err = h.reportUsecase.RejectReport(uint(id), req, superAdminID)
	if err != nil {
		utils.BuildErrorResponse(c, http.StatusInternalServerError, "Failed to reject report", err.Error())
		return
	}

	utils.BuildResponse(c, http.StatusOK, "Success", nil)
}
