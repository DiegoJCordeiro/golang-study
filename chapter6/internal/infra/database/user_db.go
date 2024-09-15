package database

import (
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/entity"
	"gorm.io/gorm"
)

type UserDB interface {
	Create(user *entity.User) error
	GetByName(username *string) (*entity.User, error)
}

type UserDBImpl struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDBImpl {
	return &UserDBImpl{
		DB: db,
	}
}

func (db *UserDBImpl) Create(user *entity.User) error {

	return db.DB.Create(user).Error
}

func (db *UserDBImpl) GetByName(username *string) (*entity.User, error) {

	var user entity.User
	err := db.DB.Where("name = ? ", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
