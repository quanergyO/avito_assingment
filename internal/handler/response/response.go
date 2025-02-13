package response

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
