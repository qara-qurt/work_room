package service

import (
	"auth/internal/model"
	"auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	repo repository.IUserRepository
}

func NewUser(repo repository.IUserRepository) *User {
	return &User{
		repo: repo,
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
