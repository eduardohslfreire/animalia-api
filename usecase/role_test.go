package usecase_test

import (
	"testing"

	"github.com/eduardohslfreire/animalia-api/entity"
	"github.com/eduardohslfreire/animalia-api/infrastructure/repository/mock"
	"github.com/eduardohslfreire/animalia-api/usecase"
	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
)

func TestFindByID(t *testing.T) {
	mockRoleRepository := new(mock.RoleRepository)
	roleUsecase := usecase.NewRoleUsecase(mockRoleRepository)

	t.Run("success - find role by id", func(t *testing.T) {
		mockRole := &entity.Role{ID: 1, Name: "Civil", Single: false}
		mockRoleRepository.On("FindByID", mocks.AnythingOfType("uint")).Return(mockRole, nil)

		role, err := roleUsecase.FindByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, role)
	})
}

func TestFindAll(t *testing.T) {
	mockRoleRepository := new(mock.RoleRepository)
	roleUsecase := usecase.NewRoleUsecase(mockRoleRepository)

	t.Run("success - find all roles", func(t *testing.T) {
		mockRoles := &entity.Roles{entity.Role{ID: 1, Name: "Civil", Single: false}}
		mockRoleRepository.On("FindAll").Return(mockRoles, nil)

		role, err := roleUsecase.FindAll()
		assert.NoError(t, err)
		assert.NotNil(t, role)
		assert.Equal(t, 1, len(*role))
	})
}

func TestFindAllCitizensByID(t *testing.T) {
	mockRoleRepository := new(mock.RoleRepository)
	roleUsecase := usecase.NewRoleUsecase(mockRoleRepository)

	t.Run("success - find all citizens by role id", func(t *testing.T) {
		mockCitizens := &entity.Citizens{entity.Citizen{ID: 1, Name: "Name", Species: "Species", Description: "Description", PhotoURL: "PhotoURL", Weight: 0, Height: 0, HasPetHuman: true}}
		mockRoleRepository.On("FindAllCitizensByID", mocks.AnythingOfType("uint")).Return(mockCitizens, nil)

		citizens, err := roleUsecase.FindAllCitizensByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, citizens)
		assert.Equal(t, 1, len(*citizens))
	})
}
