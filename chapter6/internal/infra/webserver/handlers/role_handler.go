package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type RoleHandler interface {
	CreateRole(fiberContext *fiber.Ctx) error
	UpdateRole(fiberContext *fiber.Ctx) error
	DeleteRole(fiberContext *fiber.Ctx) error
	GetAllRole(fiberContext *fiber.Ctx) error
	GetRoleByName(fiberContext *fiber.Ctx) error
}

type RoleHandlerImpl struct {
	Service service.RoleService
}

func NewRoleHandler(service service.RoleService) *RoleHandlerImpl {
	return &RoleHandlerImpl{
		Service: service,
	}
}

// CreateRole Create an Role godoc
//
// @Summary     Create Role
// @Description This endpoint is used to create an Role.
// @Tags        Role
// @Accept      json
// @Produces    json
// @Param       request      body   	dto.RoleDto      true      "Role Request"
// @Success     201
// @Failure     500         {object}    dto.Error
// @Router      /role  [post]
// @Security 	ApiKeyAuth
func (roleHandler *RoleHandlerImpl) CreateRole(fiberContext *fiber.Ctx) error {

	var roleDto dto.RoleDto

	body := bytes.NewReader(fiberContext.Body())
	errDecode := json.NewDecoder(body).Decode(&roleDto)

	if errDecode != nil {
		return errDecode
	}

	var errCreatedRole = roleHandler.Service.CreateRole(&roleDto)

	if errCreatedRole != nil {
		return errCreatedRole
	}

	fiberContext.Response().SetStatusCode(http.StatusCreated)

	return nil
}

// UpdateRole Update an Role godoc
//
// @Summary     Update Role
// @Description This endpoint is used to update an Role.
// @Tags        Role
// @Accept      json
// @Produces    json
// @Param       request      body      dto.RoleDto      true      "Role Request"
// @Success     204
// @Failure     500         {object}      dto.Error
// @Router      /role  [put]
// @Security 	ApiKeyAuth
func (roleHandler *RoleHandlerImpl) UpdateRole(fiberContext *fiber.Ctx) error {

	var roleDto dto.RoleDto

	body := bytes.NewReader(fiberContext.Body())
	errDecode := json.NewDecoder(body).Decode(&roleDto)

	if errDecode != nil {
		return errDecode
	}

	var errUpdatedRole = roleHandler.Service.UpdateRole(&roleDto)

	if errUpdatedRole != nil {
		return errUpdatedRole
	}

	fiberContext.Response().SetStatusCode(http.StatusNoContent)

	return nil
}

// DeleteRole Delete an Role godoc
//
// @Summary     Delete Role
// @Description This endpoint is used to delete an Role.
// @Tags        Role
// @Accept      json
// @Produces    json
// @Param       request      body      dto.RoleDto      true      "Role Request"
// @Success     200
// @Failure     500         {object}      dto.Error
// @Router      /role  [delete]
// @Security 	ApiKeyAuth
func (roleHandler *RoleHandlerImpl) DeleteRole(fiberContext *fiber.Ctx) error {

	var roleDto dto.RoleDto

	body := bytes.NewReader(fiberContext.Body())
	errDecode := json.NewDecoder(body).Decode(&roleDto)

	if errDecode != nil {
		return errDecode
	}

	var errDeleteRole = roleHandler.Service.DeleteRole(&roleDto)

	if errDeleteRole != nil {
		return errDeleteRole
	}

	fiberContext.Response().SetStatusCode(http.StatusOK)

	return nil
}

// GetAllRole Get All Roles godoc
//
// @Summary     Get All Roles
// @Description This endpoint is used to Get All Roles.
// @Tags        Role
// @Accept      json
// @Produces    json
// @Param       page    query    string       true      "Page Number"
// @Param       limit   query    string       true      "Limit Items on Page"
// @Param       sort    query    string       true      "Sort Page By asc or desc"
// @Success     200 	{array}  	  dto.RoleDto
// @Failure     404     {object}      dto.Error
// @Failure     500     {object}      dto.Error
// @Router      /role   [get]
// @Security 	ApiKeyAuth
func (roleHandler *RoleHandlerImpl) GetAllRole(fiberContext *fiber.Ctx) error {

	pageParam := fiberContext.Query("page")
	limitParam := fiberContext.Query("limit")

	if pageParam == "" || limitParam == "" {
		return fiber.ErrBadRequest
	}

	page, _ := strconv.Atoi(pageParam)
	limit, _ := strconv.Atoi(limitParam)

	rolesDto, errGetAllRole := roleHandler.Service.GetAllRole(page, limit, fiberContext.Query("sort"))

	if errGetAllRole != nil {
		return errGetAllRole
	}

	body, errorMarshal := json.Marshal(&rolesDto)

	if errorMarshal != nil {
		return errorMarshal
	}

	fiberContext.Response().SetStatusCode(http.StatusOK)
	_, errResponse := fiberContext.Response().BodyWriter().Write(body)

	if errResponse != nil {
		return errResponse
	}

	return nil
}

// GetRoleByName Get By Role Name godoc
//
// @Summary     Get By Role Name
// @Description This endpoint is used to Get By Role Name.
// @Tags        Role
// @Accept      json
// @Produces    json
// @Param       name    query    string       true      "Role Name"
// @Success     200 	{object}  	  dto.RoleDto
// @Failure     404     {object}      dto.Error
// @Failure     500     {object}      dto.Error
// @Router      /role/findByName  [get]
// @Security 	ApiKeyAuth
func (roleHandler *RoleHandlerImpl) GetRoleByName(fiberContext *fiber.Ctx) error {

	roleName := fiberContext.Query("name")
	roleDto, errorGetRoleByName := roleHandler.Service.GetRoleByName(&roleName)

	if errorGetRoleByName != nil {
		return errorGetRoleByName
	}

	fiberContext.Response().SetStatusCode(http.StatusOK)
	errorConvertJSON := fiberContext.JSON(roleDto)

	if errorConvertJSON != nil && roleDto == nil {
		return errorConvertJSON
	}

	return nil
}
