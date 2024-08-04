package roles

import (
	"gorm.io/gorm"
)

type Roles struct {
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Description string
	gorm.Model
}

func NewRole(name string, description string) *Roles {
	return &Roles{
		Name:        name,
		Description: description,
	}
}
