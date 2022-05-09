package usecase

import (
	"github.com/eduardohslfreire/animalia-api/entity"
	"github.com/eduardohslfreire/animalia-api/entity/dto"
)

// ICitizenUsecase ...
type ICitizenUsecase interface {
	Find(uint) (*entity.Citizen, error)
	FindAllByFilter(map[string]interface{}, dto.Pagination) (*dto.Pagination, error)
	Create(*entity.Citizen) (*entity.Citizen, error)
	Update(uint, *entity.Citizen) (*entity.Citizen, error)
	Delete(uint) error
	FindAllRolesByID(uint) (*entity.Roles, error)
	AssociateRole(uint, uint) error
	DisassociateRole(uint, uint) error
}

// IRoleUsecase ...
type IRoleUsecase interface {
	FindByID(uint) (*entity.Role, error)
	FindAll() (*entity.Roles, error)
	FindAllCitizensByID(uint) (*entity.Citizens, error)
}
