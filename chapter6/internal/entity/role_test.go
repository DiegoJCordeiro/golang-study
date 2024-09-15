package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRole_New(t *testing.T) {

	role, err := NewRole("Role Test", "Role Test")

	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, role)
}
