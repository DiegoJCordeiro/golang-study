package people

import (
	"github.com/DiegoJCordeiro/golang-study/chapter5/internal/domains/roles"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Create_People(t *testing.T) {

	sliceRoles := make([]roles.Roles, 1)
	firstRole := roles.NewRole("Test", "Test")
	sliceRoles[0] = *firstRole
	people, _ := NewPeople(0, "Test", "Test", "Test", sliceRoles)

	assert.NotNil(t, people)
}

func Test_ErrorToCreate_People(t *testing.T) {

	_, err := NewPeople(0, "", "", "", nil)

	assert.Error(t, err, "roles not must be nil or empty")
}

func FuzzTest_CreateWithID_People(f *testing.F) {

	sliceRoles := make([]roles.Roles, 1)
	firstRole := roles.NewRole("Test", "Test")
	sliceRoles[0] = *firstRole

	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for _, id := range ids {
		f.Add(id)
	}

	f.Fuzz(func(t *testing.T, id int) {
		if id == 0 {
			_, err := NewPeople(id, "Test", "Test", "Test", sliceRoles)
			assert.Error(t, err, "id cannot be 0")
		} else {
			person, _ := NewPeople(id, "Test", "Test", "Test", sliceRoles)
			assert.NotNil(t, person)
		}
	})
}
