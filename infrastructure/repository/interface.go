package repository

import (
	"github.com/eduardohslfreire/animalia-api/entity"
	"github.com/eduardohslfreire/animalia-api/entity/dto"
	"github.com/go-redsync/redsync/v4"
)

// ICitizenRepository ...
type ICitizenRepository interface {
	Create(*entity.Citizen) (*entity.Citizen, error)
	FindByID(id uint) (*entity.Citizen, error)
	FindByFilter(map[string]interface{}, dto.Pagination) (*dto.Pagination, error)
	Update(*entity.Citizen) (*entity.Citizen, error)
	Delete(*entity.Citizen) error
	FindAllRolesByID(id uint) (*entity.Roles, error)
	AddRole(*entity.Citizen, *entity.Role) error
	DeleteRole(*entity.Citizen, *entity.Role) error
}

// IRoleRepository ...
type IRoleRepository interface {
	FindByID(uint) (*entity.Role, error)
	FindAll() (*entity.Roles, error)
	FindAllCitizensByID(uint) (*entity.Citizens, error)
	CountAssociations(roleID uint) int64
}

// IRedisRepository ...
type IRedisRepository interface {
	GetValue(string) (string, bool)
	SetValue(string, interface{}, int)
	DeleteValue(key string)
	Lock(string) (*redsync.Mutex, error)
	Unlock(*redsync.Mutex) error
}
