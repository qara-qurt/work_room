package handler

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := GetTokenFromRequest(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		userId, err := ParseToken(token, []byte(h.cfg.Server.HMACSecret))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		c.Set("id", userId)
		return next(c)
	}
}

func GetTokenFromRequest(c echo.Context) (string, error) {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		return "", errors.New("authorization header is empty")
	}

	headerParts := strings.Split(header, " ")
	if headerParts[0] != "Bearer" || len(headerParts) != 2 {
		return "", errors.New("invalid authorization header")
	}
	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}

func ParseToken(token string, secret []byte) (uint, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpecting signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		return 0, err
	}

	if !t.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid subject")
	}

	id, err := strconv.Atoi(subject)
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	return uint(id), nil
}
