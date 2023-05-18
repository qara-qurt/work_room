package handler

import (
	"auth/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) SignUp(c echo.Context) error {
	var user model.UserInput
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := h.service.User.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, id)
}

func (h *Handler) SignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, "test")
}
