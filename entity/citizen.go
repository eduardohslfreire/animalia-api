package entity

import (
	"encoding/json"
	"fmt"
)

// Citizen ...
type Citizen struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Species     string
	Description string
	PhotoURL    string
	Weight      float64
	Height      float64
	HasPetHuman bool
	Roles       Roles `gorm:"many2many:citizen_role;"`
}

// Citizens ...
type Citizens []Citizen

// NewCitizen ...
func NewCitizen(name, species, description, photoURL string, weight, height float64, hasPetHuman bool) *Citizen {
	return &Citizen{
		Name:        name,
		Description: description,
		PhotoURL:    photoURL,
		Weight:      weight,
		Height:      height,
		HasPetHuman: hasPetHuman,
	}
}

// NewCitizenFromJSON ...
func NewCitizenFromJSON(citizenAsJSON string) (*Citizen, error) {
	citizen := &Citizen{}

	if err := json.Unmarshal([]byte(citizenAsJSON), citizen); err != nil {
		return nil, err
	}
	return citizen, nil
}

// AddRole ...
func (c *Citizen) AddRole(role Role) {
	c.Roles = append(c.Roles, role)
}

// GetKey ...
func (c *Citizen) GetKey() string {
	return fmt.Sprintf("CITIZEN-%d", c.ID)
}

// ToJSON ...
func (c *Citizen) ToJSON() (string, error) {
	citizenAsJSON, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(citizenAsJSON), nil
}
