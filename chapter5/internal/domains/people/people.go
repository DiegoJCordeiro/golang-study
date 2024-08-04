package people

import (
	"errors"
	"github.com/DiegoJCordeiro/golang-study/chapter5/internal/domains/roles"
	"gorm.io/gorm"
)

type People struct {
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	FirstName   string
	LastName    string
	Description string
	Roles       []roles.Roles
	gorm.Model
}

func NewPeople(id int, firstName string, lastName string, description string, roles []roles.Roles) (*People, error) {

	if id == 0 {
		return nil, errors.New("id cannot be 0")
	}

	if roles == nil || len(roles) == 0 {
		return nil, errors.New("roles not must be nil or empty")
	}

	return &People{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Description: description,
		Roles:       roles,
	}, nil
}
