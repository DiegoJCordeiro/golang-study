package service

import (
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/dto"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/entity"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/infra/database"
)

type RoleService interface {
	CreateRole(role *dto.RoleDto) error
	UpdateRole(role *dto.RoleDto) error
	DeleteRole(roleDto *dto.RoleDto) error
	GetAllRole(page, limit int, sort string) (*[]dto.RoleDto, error)
	GetRoleByName(name *string) (*dto.RoleDto, error)
}

type RoleServiceImpl struct {
	DB database.RoleDB
}

func NewRoleService(db database.RoleDB) *RoleServiceImpl {
	return &RoleServiceImpl{
		DB: db,
	}

}

func (roleService *RoleServiceImpl) CreateRole(role *dto.RoleDto) error {

	roleEntity, _ := entity.NewRole(role.Name, role.Description)
	return roleService.DB.Create(roleEntity)
}

func (roleService *RoleServiceImpl) UpdateRole(role *dto.RoleDto) error {

	_, errorGetRoleByName := roleService.GetRoleByName(&role.Name)

	if errorGetRoleByName != nil {
		return errorGetRoleByName
	}

	roleEntity, _ := entity.NewRole(role.Name, role.Description)

	return roleService.DB.Update(roleEntity)
}

func (roleService *RoleServiceImpl) DeleteRole(roleDto *dto.RoleDto) error {

	roleEntity, errorMountRoleEntity := entity.NewRole(roleDto.Name, roleDto.Description)

	if errorMountRoleEntity != nil {
		return errorMountRoleEntity
	}

	return roleService.DB.Delete(roleEntity)
}

func (roleService *RoleServiceImpl) GetAllRole(page, limit int, sort string) (*[]dto.RoleDto, error) {

	roles, errGetAll := roleService.DB.GetAll(page, limit, sort)
	rolesDto := make([]dto.RoleDto, 0)
	for _, value := range roles {
		rolesDto = append(rolesDto, *dto.NewRoleDto(value.ID, value.Name, value.Description))
	}
	return &rolesDto, errGetAll
}

func (roleService *RoleServiceImpl) GetRoleByName(name *string) (*dto.RoleDto, error) {

	role, errorGetByName := roleService.DB.GetByName(name)

	if errorGetByName != nil {
		return nil, errorGetByName
	}

	roleDto := dto.NewRoleDto(role.ID, role.Name, role.Description)

	return roleDto, nil
}
