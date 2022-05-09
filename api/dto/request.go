package dto

// PaginationRequest ...
type PaginationRequest struct {
	Sort  string `form:"sort"`
	Page  int    `form:"page" binding:"numeric,gte=0"`
	Limit int    `form:"limit" binding:"numeric,gte=0"`
}

// CitizenIDRequest ...
type CitizenIDRequest struct {
	ID uint `uri:"citizen_id" binding:"numeric,required"`
}

// RoleIDRequest ...
type RoleIDRequest struct {
	ID uint `uri:"role_id" binding:"numeric,required"`
}

// CitizenRoleIDRequest ...
type CitizenRoleIDRequest struct {
	CitizenID uint `uri:"citizen_id" binding:"numeric,required"`
	RoleID    uint `uri:"role_id" binding:"numeric,required"`
}

// FindAllCitizensQueryParamsRequest ...
type FindAllCitizensQueryParamsRequest struct {
	Name        string `form:"name"`
	Species     string `form:"species"`
	Description string `form:"description"`
	HasPetHuman bool   `form:"has_pet_human"`
	PaginationRequest
}

// CreateCitizenBodyRequest ...
type CreateCitizenBodyRequest struct {
	Name        string  `json:"name" binding:"required"`
	Species     string  `json:"species" binding:"required"`
	Description string  `json:"description" binding:"required"`
	PhotoURL    string  `json:"photo_url" binding:"required,url"`
	Weight      float64 `json:"weight" binding:"required,numeric,gt=0"`
	Height      float64 `json:"height" binding:"required,numeric,gt=0"`
	HasPetHuman bool    `json:"has_pet_human"`
}

// UpdateCitizenBodyRequest ...
type UpdateCitizenBodyRequest struct {
	Name        string  `json:"name"`
	Species     string  `json:"species"`
	Description string  `json:"description"`
	PhotoURL    string  `json:"photo_url" binding:"url"`
	Weight      float64 `json:"weight" binding:"numeric,gt=0"`
	Height      float64 `json:"height" binding:"numeric,gt=0"`
	HasPetHuman bool    `json:"has_pet_human"`
}
