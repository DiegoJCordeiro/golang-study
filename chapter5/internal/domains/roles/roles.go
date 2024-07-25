package roles

import (
	"github.com/DiegoJCordeiro/golang-study/chapter5/internal/domains/people"
	"gorm.io/gorm"
)

type Roles struct {
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Description string
	People      []people.People `gorm:"many2many:people_roles"`
	gorm.Model
}

func NewRole(name string, description string) Roles {
	return Roles{
		Name:        name,
		Description: description,
	}
}
