package usecase

import (
	"github.com/eduardohslfreire/animalia-api/entity"
	"github.com/eduardohslfreire/animalia-api/infrastructure/repository"
)

// RoleUsecase ...
type RoleUsecase struct {
	RoleRepository repository.IRoleRepository
}

// NewRoleUsecase ...
func NewRoleUsecase(citizenRepository repository.IRoleRepository) IRoleUsecase {
	return &RoleUsecase{RoleRepository: citizenRepository}
}

// FindByID ...
func (r *RoleUsecase) FindByID(id uint) (*entity.Role, error) {
	return r.RoleRepository.FindByID(id)
}

// FindAll ...
func (r *RoleUsecase) FindAll() (*entity.Roles, error) {
	return r.RoleRepository.FindAll()
}

// FindAllCitizensByID ...
func (r *RoleUsecase) FindAllCitizensByID(id uint) (*entity.Citizens, error) {
	return r.RoleRepository.FindAllCitizensByID(id)
}
