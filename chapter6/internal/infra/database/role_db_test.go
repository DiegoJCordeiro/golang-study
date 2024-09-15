package database

import (
	"fmt"
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestRoleDB_Create(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	err = db.AutoMigrate(&entity.Role{})
	assert.Nil(t, err)

	role, err := entity.NewRole("Role Test", " Role Test.")
	assert.Nil(t, err)

	roleDB := NewRoleDB(db)
	errRoleUser := roleDB.Create(role)
	assert.NoError(t, errRoleUser)
}

func TestRoleDB_Update(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&entity.Role{})
	assert.NoError(t, err)

	role, err := entity.NewRole("Role Test", " Role Test.")
	assert.NoError(t, err)

	roleDB := NewRoleDB(db)
	errRoleUserCreate := roleDB.Create(role)
	assert.NoError(t, errRoleUserCreate)

	role.Description = "Role Test Updated"
	errRoleUserUpdate := roleDB.Update(role)
	assert.NoError(t, errRoleUserUpdate)

	roleFound, errRoleFound := roleDB.GetByName("Role Test")
	assert.NoError(t, errRoleFound)
	assert.Equal(t, roleFound.Description, "Role Test Updated")
}

func TestRoleDB_Delete(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&entity.Role{})
	assert.NoError(t, err)

	role, err := entity.NewRole("Role Test", " Role Test.")
	assert.NoError(t, err)

	roleDB := NewRoleDB(db)
	errRoleUserCreate := roleDB.Create(role)
	assert.NoError(t, errRoleUserCreate)

	errRoleUserDelete := roleDB.Delete(role)
	assert.NoError(t, errRoleUserDelete)

	roleFound, errGetByName := roleDB.GetByName("Role Test")
	assert.Error(t, errGetByName)
	assert.Nil(t, roleFound)
}

func TestRoleDB_GetAll(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&entity.Role{})
	assert.NoError(t, err)

	roleDB := NewRoleDB(db)

	for index := 0; index < 10; index++ {

		role, err := entity.NewRole(fmt.Sprintf("Role Test %d", index), " Role Test.")
		assert.NoError(t, err)

		errRoleUserCreate := roleDB.Create(role)
		assert.NoError(t, errRoleUserCreate)
	}

	roles, errRolesFound := roleDB.GetAll(0, 5, "asc")
	assert.NoError(t, errRolesFound)
	assert.Equal(t, "Role Test 0", roles[0].Name)
	assert.Equal(t, "Role Test 5", roles[5].Name)
}

func TestRoleDB_GetByName(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&entity.Role{})
	assert.NoError(t, err)

	role, err := entity.NewRole("Role Test", " Role Test.")
	assert.NoError(t, err)

	roleDB := NewRoleDB(db)
	errRoleUserCreate := roleDB.Create(role)
	assert.NoError(t, errRoleUserCreate)

	roleFound, errRoleFound := roleDB.GetByName("Role Test")
	assert.NoError(t, errRoleFound)
	assert.NotNil(t, roleFound)
}
