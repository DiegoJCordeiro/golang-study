package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type ErrorAPI struct {
}

func NewErrorAPI() *ErrorAPI {
	return &ErrorAPI{}
}

func (handlerErrorAPI *ErrorAPI) ErrorHandler(ctx *fiber.Ctx, err error) error {

	var fiberError *fiber.Error

	if errors.As(err, &fiberError) {
		return fiberError
	}

	return ctx.Status(handlerErrorAPI.checkStatusCode(err)).SendString(err.Error())
}

func (handlerErrorAPI *ErrorAPI) checkStatusCode(err error) int {

	if err != nil {

		errorUpperCase := strings.ToUpper(err.Error())

		if strings.Contains(errorUpperCase, strings.ToUpper("Not Found")) {
			return fiber.StatusNotFound
		}

		if strings.Contains(errorUpperCase, strings.ToUpper("Invalid")) {
			return fiber.StatusBadRequest
		}

		if strings.Contains(errorUpperCase, strings.ToUpper("Unauthorized")) {
			return fiber.StatusUnauthorized
		}

	}
	return fiber.StatusInternalServerError
}
