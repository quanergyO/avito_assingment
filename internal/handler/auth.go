package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanergyO/avito_assingment/internal/handler/response"
	"github.com/quanergyO/avito_assingment/types"
)

func (h *Handler) SignUp(c *gin.Context) {
	const op = "Handler.SignUp"

	log := slog.With(
		slog.String("op", op),
	)

	log.Info("Call /SignUp")
	var user types.UserType
	if err := c.BindJSON(&user); err != nil {
		log.Error("Invalid input body")
		response.NewErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input types.SignInInput

	if err := c.BindJSON(&input); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.CheckAuthData(input.Username, input.Password)
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := h.service.GenerateToken(user)
	if err != nil {
		response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
