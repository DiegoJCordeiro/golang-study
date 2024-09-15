package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_New(t *testing.T) {

	user, err := NewUser("Test Test", "test@test.com", "123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestUser_ValidatePassword(t *testing.T) {

	user, err := NewUser("Test Test", "test@test.com", "123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123"))
}
