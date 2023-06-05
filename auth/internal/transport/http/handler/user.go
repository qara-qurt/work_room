package handler

import (
	"auth/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) SignUp(c echo.Context) error {
	var user model.UserInput
	if err := c.Bind(&user); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if err := user.Validate(); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	id, err := h.service.User.Create(user)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]int{
		"id": id,
	})
}

func (h *Handler) SignIn(c echo.Context) error {
	var user model.UserAuthInput
	if err := c.Bind(&user); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := user.Validate(); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	accessToken, refreshToken, err := h.service.User.SignIn(user)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
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
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
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

func (h *Handler) GetUsers(c echo.Context) error {
	pageStr := c.QueryParam("page")
	offsetStr := c.QueryParam("offset")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 25
	}

	users, err := h.service.User.GetUsers(page, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string][]model.User{
		"users": users,
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	user, err := h.service.User.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}
