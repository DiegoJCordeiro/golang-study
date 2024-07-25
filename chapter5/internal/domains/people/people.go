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
	Roles       []roles.Roles `gorm:"many2many:people_roles"`
	gorm.Model
}

func NewPeople(firstName string, lastName string, description string, roles []roles.Roles) (*People, error) {

	if roles != nil && len(roles) > 0 {

		return nil, errors.New("roles not must be nil or empty")
	}

	return &People{
		FirstName:   firstName,
		LastName:    lastName,
		Description: description,
		Roles:       roles,
	}, nil
}
