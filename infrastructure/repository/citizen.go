package repository

import (
	"math"

	"github.com/eduardohslfreire/animalia-api/entity"
	"github.com/eduardohslfreire/animalia-api/entity/dto"
	"gorm.io/gorm"
)

// CitizenRepository ...
type CitizenRepository struct {
	Conn *gorm.DB
}

// NewCitizenRepository ...
func NewCitizenRepository(Conn *gorm.DB) ICitizenRepository {
	return &CitizenRepository{Conn}
}

// Create ...
func (c *CitizenRepository) Create(citizen *entity.Citizen) (*entity.Citizen, error) {
	if err := c.Conn.Create(citizen).Error; err != nil {
		return nil, err
	}

	return citizen, nil
}

// FindByID ...
func (c *CitizenRepository) FindByID(id uint) (*entity.Citizen, error) {
	citizen := &entity.Citizen{ID: id}
	if err := c.Conn.Where("id = $1 ", id).Preload("Roles").First(citizen).Error; err != nil {
		return nil, err
	}

	return citizen, nil
}

// FindByFilter ...
func (c *CitizenRepository) FindByFilter(filters map[string]interface{}, pagination dto.Pagination) (*dto.Pagination, error) {
	citizens := &entity.Citizens{}

	if err := c.Conn.Scopes(paginate(filters, &pagination, c.Conn)).Find(citizens).Error; err != nil {
		return nil, err
	}
	pagination.Rows = citizens

	return &pagination, nil
}

// Update ...
func (c *CitizenRepository) Update(citizen *entity.Citizen) (*entity.Citizen, error) {
	if err := c.Conn.Save(citizen).Error; err != nil {
		return nil, err
	}

	return citizen, nil
}

// Delete ...
func (c *CitizenRepository) Delete(citizen *entity.Citizen) error {
	if err := c.Conn.Select("Roles").Delete(citizen).Error; err != nil {
		return err
	}
	return nil
}

// FindAllRolesByID ...
func (c *CitizenRepository) FindAllRolesByID(id uint) (*entity.Roles, error) {
	citizen := &entity.Citizen{ID: id}
	roles := &entity.Roles{}
	if err := c.Conn.Model(citizen).Association("Roles").Find(roles); err != nil {
		return nil, err
	}
	return roles, nil
}

// AddRole ...
func (c *CitizenRepository) AddRole(citizen *entity.Citizen, role *entity.Role) error {
	targetCitizen := &entity.Citizen{ID: citizen.ID}
	if err := c.Conn.Model(targetCitizen).Association("Roles").Append(role); err != nil {
		return err
	}
	return nil
}

// DeleteRole ...
func (c *CitizenRepository) DeleteRole(citizen *entity.Citizen, role *entity.Role) error {
	if err := c.Conn.Model(citizen).Association("Roles").Delete(role); err != nil {
		return err
	}
	return nil
}

func paginate(filters map[string]interface{}, pagination *dto.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(new(entity.Citizens)).Where(filters).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(filters).Preload("Roles").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
