package entity

import (
	"errors"
	"github.com/DiegoJCordeiro/golang-study/chapter6/pkg/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID `gorm:"primaryKey;"`
	Name        string    `gorm:"unique;not null"`
	Description string
	gorm.Model
}

func NewRole(name string, description string) (*Role, error) {

	if name == "" || description == "" {
		return nil, errors.New("name and description are required")
	}

	return &Role{
		ID:          entity.NewID(),
		Name:        name,
		Description: description,
	}, nil
}
