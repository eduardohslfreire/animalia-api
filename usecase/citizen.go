package usecase

import (
	"fmt"

	"github.com/eduardohslfreire/animalia-api/config/env"
	"github.com/eduardohslfreire/animalia-api/entity"
	"github.com/eduardohslfreire/animalia-api/entity/dto"
	"github.com/eduardohslfreire/animalia-api/entity/errors"
	"github.com/eduardohslfreire/animalia-api/infrastructure/repository"
	"github.com/jinzhu/copier"
)

// CitizenUsecase ...
type CitizenUsecase struct {
	CitizenRepository repository.ICitizenRepository
	RoleRepository    repository.IRoleRepository
	RedisRepository   repository.IRedisRepository
}

var citizenRedisKey = func(id uint) string { return fmt.Sprintf("CITIZEN-%d", id) }

// NewCitizenUsecase ...
func NewCitizenUsecase(citizenRepository repository.ICitizenRepository, roleRepository repository.IRoleRepository, redisRepository repository.IRedisRepository) ICitizenUsecase {
	return &CitizenUsecase{CitizenRepository: citizenRepository, RoleRepository: roleRepository, RedisRepository: redisRepository}
}

// Find ...
func (c *CitizenUsecase) Find(id uint) (*entity.Citizen, error) {
	if citizenAsJSON, has := c.RedisRepository.GetValue(citizenRedisKey(id)); has {
		if citizen, err := entity.NewCitizenFromJSON(citizenAsJSON); err == nil {
			return citizen, nil
		}
	}

	citizen, err := c.CitizenRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	go c.saveOrUpdateCache(citizen)

	return citizen, nil
}

// FindAllByFilter ...
func (c *CitizenUsecase) FindAllByFilter(filters map[string]interface{}, pagination dto.Pagination) (*dto.Pagination, error) {
	return c.CitizenRepository.FindByFilter(filters, pagination)
}

// Create ...
func (c *CitizenUsecase) Create(citizen *entity.Citizen) (*entity.Citizen, error) {
	citizen, err := c.CitizenRepository.Create(citizen)
	if err != nil {
		return nil, err
	}

	go c.saveOrUpdateCache(citizen)

	return citizen, nil
}

// Update ...
func (c *CitizenUsecase) Update(id uint, citizen *entity.Citizen) (*entity.Citizen, error) {
	targetCitizen, err := c.Find(id)
	if err != nil {
		return nil, err
	}
	citizen.ID = id
	copier.Copy(targetCitizen, citizen)

	targetCitizen, err = c.CitizenRepository.Update(targetCitizen)
	if err != nil {
		return nil, err
	}

	go c.saveOrUpdateCache(targetCitizen)

	return targetCitizen, nil
}

// Delete ...
func (c *CitizenUsecase) Delete(id uint) error {
	citizen, err := c.Find(id)
	if err != nil {
		return err
	}

	if err := c.CitizenRepository.Delete(citizen); err != nil {
		return err
	}

	go c.RedisRepository.DeleteValue(citizenRedisKey(id))

	return nil
}

// FindAllRolesByID ...
func (c *CitizenUsecase) FindAllRolesByID(id uint) (*entity.Roles, error) {
	return c.CitizenRepository.FindAllRolesByID(id)
}

// AssociateRole ...
func (c *CitizenUsecase) AssociateRole(citizenID uint, roleID uint) (err error) {
	defer c.reloadCacheIfSuccess(err, citizenID)

	citizen, err := c.CitizenRepository.FindByID(citizenID)
	if err != nil {
		return err
	}

	role, err := c.RoleRepository.FindByID(roleID)
	if err != nil {
		return err
	}

	citizenRoles := citizen.Roles
	if citizenRoles != nil && len(citizenRoles) > 0 {
		if citizenRoles.HasRole(role.Name) {
			return nil
		}

		if role.IsCivilRole() {
			return errors.NewBusinessError("Civil role cannot be assigned to a citizen that already has another role")
		}

		if citizenRoles.HasCivilRole() {
			return errors.NewBusinessError("Citizen already performs the role of civil and cannot perform another role")
		}
	}

	if role.IsSingle() {
		mutex, err := c.RedisRepository.Lock(role.GetKey())
		if err != nil {
			return err
		}
		defer c.RedisRepository.Unlock(mutex)

		if c.existsAnyCitizenWithRole(roleID) {
			return errors.NewBusinessError("Role is already occupied and can only be performed by a citizen")
		}
	}

	return c.CitizenRepository.AddRole(citizen, role)
}

// DisassociateRole ...
func (c *CitizenUsecase) DisassociateRole(citizenID uint, roleID uint) (err error) {
	defer c.reloadCacheIfSuccess(err, citizenID)

	citizen, err := c.CitizenRepository.FindByID(citizenID)
	if err != nil {
		return err
	}

	role, err := c.RoleRepository.FindByID(roleID)
	if err != nil {
		return err
	}

	return c.CitizenRepository.DeleteRole(citizen, role)
}

func (c *CitizenUsecase) existsAnyCitizenWithRole(roleID uint) bool {
	return c.RoleRepository.CountAssociations(roleID) > int64(0)
}

func (c *CitizenUsecase) saveOrUpdateCache(citizen *entity.Citizen) {
	if citizenAsJSON, err := citizen.ToJSON(); err == nil {
		c.RedisRepository.SetValue(citizen.GetKey(), citizenAsJSON, env.RedisExpirationHours)
	}
}

func (c *CitizenUsecase) reloadCacheIfSuccess(err error, citizenID uint) {
	if err != nil {
		return
	}

	if citizen, err := c.CitizenRepository.FindByID(citizenID); err == nil {
		c.saveOrUpdateCache(citizen)
	}
}
