package handler

import (
	"auth/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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

	return c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

func (h *Handler) SignIn(c echo.Context) error {
	var user model.UserAuthInput
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := user.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	accessToken, refreshToken, err := h.service.User.SignIn(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(2 * 30 * 24 * time.Hour), // expires 2 month
		Path:     "/",
		HttpOnly: true,
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{
		"token": accessToken,
	})
}

func (h *Handler) Refresh(c echo.Context) error {
	token, err := c.Cookie("refresh_token")
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	accessToken, refreshToken, err := h.service.User.RefreshTokens(token.Value)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(2 * 30 * 24 * time.Hour), // expires 2 month
		Path:     "/",
		HttpOnly: true,
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{
		"token": accessToken,
	})
}
