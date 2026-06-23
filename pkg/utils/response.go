package utils

import (
	"github.com/gin-gonic/gin"
	"toba-lapor-backend/internal/model/dto"
)

func BuildResponse(c *gin.Context, code int, status string, data interface{}) {
	c.JSON(code, dto.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	})
}

func BuildErrorResponse(c *gin.Context, code int, status string, err string) {
	c.AbortWithStatusJSON(code, dto.ErrorResponse{
		Code:   code,
		Status: status,
		Error:  err,
	})
}
