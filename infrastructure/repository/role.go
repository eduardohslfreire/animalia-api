package repository

import (
	"github.com/eduardohslfreire/animalia-api/entity"
	"gorm.io/gorm"
)

// RoleRepository ...
type RoleRepository struct {
	Conn *gorm.DB
}

// NewRoleRepository ...
func NewRoleRepository(Conn *gorm.DB) IRoleRepository {
	return &RoleRepository{Conn}
}

// FindByID ...
func (r *RoleRepository) FindByID(id uint) (*entity.Role, error) {
	role := &entity.Role{ID: id}
	if err := r.Conn.First(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

// FindAll ...
func (r *RoleRepository) FindAll() (*entity.Roles, error) {
	roles := &entity.Roles{}
	if err := r.Conn.Find(roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

// FindAllCitizensByID ...
func (r *RoleRepository) FindAllCitizensByID(id uint) (*entity.Citizens, error) {
	role := &entity.Role{ID: id}
	citizens := &entity.Citizens{}
	if err := r.Conn.Model(role).Association("Citizens").Find(citizens); err != nil {
		return nil, err
	}

	return citizens, nil
}

// CountAssociations ...
func (r *RoleRepository) CountAssociations(roleID uint) int64 {
	role := &entity.Role{ID: roleID}
	return r.Conn.Model(role).Association("Citizens").Count()
}
