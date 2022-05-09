package entity

import (
	"fmt"

	"github.com/eduardohslfreire/animalia-api/entity/enum"
)

// Role ...
type Role struct {
	ID       uint `gorm:"primary_key"`
	Name     enum.RoleName
	Single   bool
	Citizens []Citizen `gorm:"many2many:citizen_role;"`
}

// IsSingle ...
func (r Role) IsSingle() bool {
	return r.Single
}

// IsCivilRole ...
func (r Role) IsCivilRole() bool {
	return r.Name == enum.Civil
}

// GetKey ...
func (r Role) GetKey() string {
	return fmt.Sprintf("ROLE-%d", r.ID)
}

// Roles ...
type Roles []Role

// HasCivilRole ...
func (r *Roles) HasCivilRole() bool {
	return r.HasRole(enum.Civil)
}

// HasRole ...
func (r *Roles) HasRole(roleName enum.RoleName) bool {
	for _, role := range *r {
		if role.Name == roleName {
			return true
		}
	}
	return false
}
