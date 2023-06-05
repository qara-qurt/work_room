package service

import (
	"auth/internal/model"
	"auth/internal/repository"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	repo        repository.IUserRepository
	sessionRepo repository.ITokenRepository
	hmacSecret  []byte
}

func NewUser(repo repository.IUserRepository, sessionRepo repository.ITokenRepository, secret string) *User {
	return &User{
		repo:        repo,
		sessionRepo: sessionRepo,
		hmacSecret:  []byte(secret),
	}
}

func (u *User) Create(user model.UserInput) (int, error) {
	hashPass, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashPass
	return u.repo.Create(user)
}

func (u *User) SignIn(user model.UserAuthInput) (string, string, error) {
	// get user by email
	res, err := u.repo.GetByEmail(user.Email)
	if err != nil {
		return "", "", err
	}

	// check password is correct
	if ok := checkPasswordHash(user.Password, res.Password); !ok {
		return "", "", model.ErrorUserPassword
	}

	return u.GenerateTokens(res.ID)
}

func (u *User) GenerateTokens(userID uint) (string, string, error) {
	// generate access_token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   strconv.Itoa(int(userID)),
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	// signed with HMAC_secret
	accessToken, err := token.SignedString(u.hmacSecret)
	if err != nil {
		return "", "", err
	}

	// generate refresh_token
	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}

	// save refresh token in redis
	if err := u.sessionRepo.Create(model.RefreshSession{
		UserID: userID,
		Token:  refreshToken,
	}); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (u *User) RefreshTokens(refreshToken string) (string, string, error) {
	session, err := u.sessionRepo.Get(refreshToken)
	if err != nil {
		return "", "", err
	}

	return u.GenerateTokens(session.UserID)
}

func (u *User) ParseToken(ctx context.Context, token string) (uint, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpecting signing method: %v", token.Header["alg"])
		}

		return u.hmacSecret, nil
	})
	if err != nil {
		return 0, err
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

func (u *User) GetUsers(page, offset int) ([]model.User, error) {
	return u.repo.GetUsers(page, offset)
}

func (u *User) GetUser(id int) (model.User, error) {
	return u.repo.GetUser(id)
}

func newRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
