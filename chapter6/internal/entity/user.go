package entity

import (
	"github.com/DiegoJCordeiro/golang-study/chapter6/pkg/entity"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Password string
}

func NewUser(name, email, password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (user *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
