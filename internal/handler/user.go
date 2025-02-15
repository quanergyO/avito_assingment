package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanergyO/avito_assingment/internal/handler/response"
)

func (h *Handler) GetInfo(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userInfo, err := h.service.GetUserInfo(userId)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": userInfo,
	})

}

func (h *Handler) SendCoin(c *gin.Context) {
	_, err := h.getUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Not implemented",
	})
}

func (h *Handler) BuyItem(c *gin.Context) {
	_, err := h.getUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Not implemented",
	})
}
