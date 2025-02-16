package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quanergyO/avito_assingment/internal/handler/response"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		response.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		response.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	claims, err := h.service.ParserToken(headerParts[1])
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("UserId", claims.UserId)
}

func (h *Handler) getUserId(c *gin.Context) (int, error) {
	userIdIface, exists := c.Get("UserId")
	userId, ok := userIdIface.(int)
	if !ok {
		return 0, fmt.Errorf("can't convert userid")
	}
	if !exists || userId == 0 {
		return 0, fmt.Errorf("bad token claims")
	}

	return userId, nil
}
