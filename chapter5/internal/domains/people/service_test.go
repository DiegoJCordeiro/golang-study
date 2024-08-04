package people

import (
	"github.com/DiegoJCordeiro/golang-study/chapter5/internal/domains/roles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type RepositoryMock struct {
	mock.Mock
}

func (repositoryMock *RepositoryMock) GetPeople(id int) (*People, error) {
	args := repositoryMock.Called(id)
	return args.Get(0).(*People), args.Error(1)
}

func Test_Service_GetPeople(t *testing.T) {

	sliceRoles := make([]roles.Roles, 1)
	firstRole := roles.NewRole("Test", "Test")
	sliceRoles[0] = *firstRole
	people, _ := NewPeople(1, "Test", "Test", "Test", sliceRoles)

	assertPeople := assert.New(t)
	assertPeople.NotNil(people)

	repositoryMock := new(RepositoryMock)
	repositoryMock.On("GetPeople", 1).Return(people, nil)

	service := Service{
		Repository: repositoryMock,
	}

	peopleFound, _ := service.Repository.GetPeople(1)
	
	assertPeople.NotNil(peopleFound)
	repositoryMock.AssertExpectations(t)
}
