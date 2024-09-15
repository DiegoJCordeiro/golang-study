package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler interface {
	CreateUser(fiberContext *fiber.Ctx) error
	GetUserByName(fiberContext *fiber.Ctx) error
}

type UserHandlerImpl struct {
	Service service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{
		Service: userService,
	}
}

// CreateUser Create an User godoc
//
// @Summary     Create User
// @Description This endpoint is used to create an User.
// @Tags        User
// @Accept      json
// @Produces    json
// @Param       request      body      dto.UserDto      true      "User Request"
// @Success     201
// @Failure     500         {object}      dto.Error
// @Router      /user  [post]
func (userHandler *UserHandlerImpl) CreateUser(fiberContext *fiber.Ctx) error {

	var userDto dto.UserDto

	body := bytes.NewReader(fiberContext.Body())
	errDecode := json.NewDecoder(body).Decode(&userDto)

	if errDecode != nil {
		return fiber.ErrBadRequest
	}

	errorToCreate := userHandler.Service.Create(&userDto)

	if errorToCreate != nil {
		return fiber.ErrInternalServerError
	}

	return nil
}

// GetUserByName Get an User  By Name godoc
// @Summary     Get an User
// @Description This endpoint is used to Get an User.
// @Tags        User
// @Accept      json
// @Produces    json
// @Param       username    query    string       false      "User name"
// @Success     200 		{object}  	  dto.UserDto
// @Failure     404         {object}      dto.Error
// @Failure     500         {object}      dto.Error
// @Router      /user/findByName  [get]
// @Security 	ApiKeyAuth
func (userHandler *UserHandlerImpl) GetUserByName(fiberContext *fiber.Ctx) error {

	username := fiberContext.Query("username")

	userDto, errGetUserByName := userHandler.Service.GetUserByName(&username)

	if errGetUserByName != nil {
		return fiber.ErrNotFound
	}

	errorConvertJson := fiberContext.JSON(userDto)

	if errorConvertJson != nil {
		return fiber.ErrInternalServerError
	}

	fiberContext.Response().SetStatusCode(http.StatusOK)

	return nil
}
