package database

import (
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/entity"
	"gorm.io/gorm"
)

type RoleDB interface {
	Create(role *entity.Role) error
	Update(role *entity.Role) error
	Delete(role *entity.Role) error
	GetAll(page, limit int, sort string) ([]entity.Role, error)
	GetByName(name *string) (*entity.Role, error)
}

type RoleDBImpl struct {
	DB gorm.DB
}

func NewRoleDB(db *gorm.DB) *RoleDBImpl {
	return &RoleDBImpl{DB: *db}
}

func (roleDB *RoleDBImpl) Create(role *entity.Role) error {
	return roleDB.DB.Create(role).Error
}

func (roleDB *RoleDBImpl) Update(role *entity.Role) error {

	return roleDB.DB.Where("ID = ?", role.ID).Updates(role).Error
}

func (roleDB *RoleDBImpl) Delete(role *entity.Role) error {

	roleFound, errGetByName := roleDB.GetByName(&role.Name)

	if errGetByName != nil {
		return errGetByName
	}

	return roleDB.DB.Where("name = ?", role.Name).Delete(roleFound).Error
}

func (roleDB *RoleDBImpl) GetAll(page, limit int, sort string) ([]entity.Role, error) {

	var roles []entity.Role
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page > 0 && limit > 0 {
		err = roleDB.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&roles).Error
	} else {
		err = roleDB.DB.Order("created_at " + sort).Find(&roles).Error
	}

	return roles, err
}

func (roleDB *RoleDBImpl) GetByName(name *string) (*entity.Role, error) {

	var role entity.Role
	err := roleDB.DB.Where("name = ? ", name).First(&role).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}
