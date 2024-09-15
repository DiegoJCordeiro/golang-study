package dto

import "github.com/google/uuid"

type UserDto struct {
	ID       uuid.UUID `json:"-" example:"-"`
	Name     string    `json:"name" example:"Test123"`
	Email    string    `json:"email" example:"test.test@gmail.com"`
	Password string    `json:"password" example:"Pass123"`
}

func NewUserDto(name string, email string, password string) *UserDto {

	return &UserDto{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
