package handler

import (
	"auth/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) CreateCompany(c echo.Context) error {
	var company model.CompanyInput
	if err := c.Bind(&company); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := company.Validate(); err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	id, err := h.service.Company.Create(company)
	if err != nil {
		logrus.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]int{
		"id": id,
	})
}
