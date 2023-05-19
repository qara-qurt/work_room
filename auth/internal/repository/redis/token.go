package redis

import (
	"auth/internal/model"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type Token struct {
	db *redis.Client
}

func NewToken(db *redis.Client) *Token {
	return &Token{
		db: db,
	}
}

func (t Token) Create(token model.RefreshSession) error {
	expiration := 2 * 30 * 24 * time.Hour //2 month
	return t.db.Set(context.TODO(), token.Token, fmt.Sprintf("%d", token.UserID), expiration).Err()
}

func (t Token) Get(token string) (model.RefreshSession, error) {
	userID, err := t.db.Get(context.Background(), token).Result()
	if err == redis.Nil {
		return model.RefreshSession{}, model.ErrorTokenExpired
	} else if err != nil {
		return model.RefreshSession{}, err
	}

	if err := t.db.Del(context.TODO(), token).Err(); err != nil {
		return model.RefreshSession{}, err
	}

	ID, err := strconv.Atoi(userID)
	if err != nil {
		return model.RefreshSession{}, err
	}
	return model.RefreshSession{
		UserID: uint(ID),
		Token:  token,
	}, nil

}
