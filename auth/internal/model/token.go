package model

import "errors"

type RefreshSession struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

var ErrorTokenExpired = errors.New("token is expired")
