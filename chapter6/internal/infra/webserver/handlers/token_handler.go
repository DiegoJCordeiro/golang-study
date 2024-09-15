package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/infra/environment"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenHandler interface {
	GenerateToken(fiberContext *fiber.Ctx) error
}

type TokenHandlerImpl struct {
	Service service.TokenService
}

func NewTokenHandler(service service.TokenService) *TokenHandlerImpl {
	return &TokenHandlerImpl{Service: service}
}

// GenerateToken Generate an User Token godoc
//
// @Summary     Create User Token
// @Description This endpoint is used to create an User Token.
// @Tags        Token
// @Accept      json
// @Produces    json
// @Param       request       body      	dto.UserDto      true      "User Request"
// @Success      200
// @Failure     401          {object}      	dto.Error
// @Failure     500          {object}      	dto.Error
// @Router      /user/token  [post]
func (handler *TokenHandlerImpl) GenerateToken(fiberContext *fiber.Ctx) error {

	var userDto dto.UserDto

	body := bytes.NewReader(fiberContext.Body())
	errDecode := json.NewDecoder(body).Decode(&userDto)

	if errDecode != nil {
		return fiber.ErrBadRequest
	}

	errorInvalidUser := handler.Service.ValidateUser(userDto.Name, userDto.Password)

	if errorInvalidUser != nil {
		return fiber.ErrUnauthorized
	}

	claims := jwt.MapClaims{
		"username": userDto.Name,
		"exp":      time.Now().Add(time.Second * time.Duration(environment.GetJwtExpiresIn())).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenGenerated, err := token.SignedString([]byte(environment.GetJwtSecret()))

	if err != nil {
		return fiberContext.SendStatus(fiber.StatusInternalServerError)
	}

	errorConvertJSON := fiberContext.JSON(fiber.Map{
		"token": tokenGenerated,
	})

	if errorConvertJSON != nil {
		return fiberContext.SendStatus(fiber.StatusInternalServerError)
	}

	fiberContext.Response().SetStatusCode(fiber.StatusOK)

	return nil
}
