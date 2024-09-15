package service

import (
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/entity"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/infra/database"
)

type UserService interface {
	Create(user *dto.UserDto) error
	GetUserByName(username *string) (*dto.UserDto, error)
}

type UserServiceImpl struct {
	DB database.UserDB
}

func NewUserService(db database.UserDB) *UserServiceImpl {
	return &UserServiceImpl{
		DB: db,
	}
}

func (userService *UserServiceImpl) Create(user *dto.UserDto) error {

	userEntity, errorNewUser := entity.NewUser(user.Name, user.Email, user.Password)

	if errorNewUser != nil {
		return errorNewUser
	}

	errorCreateUser := userService.DB.Create(userEntity)

	if errorCreateUser != nil {
		return errorCreateUser
	}

	return nil
}

func (userService *UserServiceImpl) GetUserByName(username *string) (*dto.UserDto, error) {

	userEntity, errorGetByName := userService.DB.GetByName(username)

	if errorGetByName != nil {
		return nil, errorGetByName
	}

	userDto := dto.NewUserDto(userEntity.Name, userEntity.Email, userEntity.Password)

	return userDto, nil
}
