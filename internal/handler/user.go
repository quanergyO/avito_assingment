package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanergyO/avito_assingment/internal/handler/response"
	"github.com/quanergyO/avito_assingment/types"
)

func (h *Handler) GetInfo(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		slog.Warn("Invalid token")
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
	userId, err := h.getUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var requestData types.SendCoinRequest
	if err := c.BindJSON(&requestData); err != nil {
		slog.Error("Invalid input body")
		response.NewErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	if err := h.service.User.SendCoins(userId, requestData.ReceiverId, requestData.Amount); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, "internal sevice error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}

func (h *Handler) BuyItem(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item := c.Param("item")
	if err := h.service.User.BuyItem(userId, item); err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}
