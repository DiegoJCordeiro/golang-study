package database

import (
	"github.com/DiegoJCordeiro/golang-study/chapter6/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestUserDB_Create(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	err = db.AutoMigrate(&entity.User{})
	assert.Nil(t, err)

	user, err := entity.NewUser("Test Test", "test@test.com", "123")
	assert.Nil(t, err)

	userDB := NewUserDB(db)
	errCreationUser := userDB.Create(user)
	assert.Nil(t, errCreationUser)
}

func TestUserDB_GetByName(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.Nil(t, err)

	err = db.AutoMigrate(&entity.User{})
	assert.Nil(t, err)

	userDB := NewUserDB(db)
	userCreated, err := entity.NewUser("Test Test", "test@test.com", "123")

	errCreationUser := userDB.Create(userCreated)
	assert.Nil(t, errCreationUser)

	userFound, err := userDB.GetByName("Test Test")

	assert.Nil(t, err)
	assert.NotNil(t, userFound)
	assert.Equal(t, userFound.ID, userCreated.ID)
	assert.Equal(t, userCreated.Name, userFound.Name)
	assert.Equal(t, userCreated.Email, userFound.Email)
	assert.True(t, userFound.ValidatePassword("123"))
}
