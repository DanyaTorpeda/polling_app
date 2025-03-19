package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func newErrorResponse(c *gin.Context, code int, message string, logger *logrus.Logger) {
	logger.Errorf("Error %d: %s", code, message)
	c.JSON(code, errorResponse{Message: message, Code: code})
}
