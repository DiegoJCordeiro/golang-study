package service

import (
	"errors"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/infra/database"
)

type TokenService interface {
	ValidateUser(username, password string) error
}

type TokenServiceImpl struct {
	UserDB database.UserDB
}

func NewTokenService(userDB database.UserDB) TokenService {
	return &TokenServiceImpl{UserDB: userDB}
}

func (tokenService *TokenServiceImpl) ValidateUser(username, password string) error {

	if username == "" && password == "" {
		return errors.New("username or Password is empty")
	}

	user, _ := tokenService.UserDB.GetByName(&username)

	if user == nil {
		return errors.New("user not found")
	}

	isPasswordValid := user.ValidatePassword(password)

	if !isPasswordValid {
		return errors.New("invalid password")
	}

	return nil
}
