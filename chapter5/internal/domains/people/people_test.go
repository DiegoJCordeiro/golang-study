package people

import (
	"github.com/DiegoJCordeiro/golang-study/chapter5/internal/domains/roles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_People(t *testing.T) {

	sliceRoles := make([]roles.Roles, 1)
	sliceRoles = append(sliceRoles, roles.NewRole("", ""))
	people, _ := NewPeople("", "", "", sliceRoles)

	assert.NotNil(t, people)
}

func Test_CreateWithError_People(t *testing.T) {

	_, err := NewPeople("", "", "", nil)

	assert.Error(t, err, "roles not must be nil or empty")
}
