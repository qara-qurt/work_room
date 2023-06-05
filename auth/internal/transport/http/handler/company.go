package handler

import (
	"auth/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

func (h *Handler) GetCompanies(c echo.Context) error {
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

	companies, err := h.service.Company.GetCompanies(page, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string][]model.Company{
		"companies": companies,
	})
}

func (h *Handler) GetCompany(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	company, err := h.service.Company.GetCompany(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, company)
}
