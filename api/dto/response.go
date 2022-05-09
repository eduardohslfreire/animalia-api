package dto

// PaginationCitizenBodyResponse ...
type PaginationCitizenBodyResponse struct {
	Limit      int                  `json:"limit,omitempty;query:limit"`
	Page       int                  `json:"page,omitempty;query:page"`
	Sort       string               `json:"sort,omitempty;query:sort"`
	TotalRows  int64                `json:"total_rows"`
	TotalPages int                  `json:"total_pages"`
	Rows       CitizensBodyResponse `json:"rows"`
}

// CitizensBodyResponse ...
type CitizensBodyResponse []CitizenBodyResponse

// CitizenBodyResponse ...
type CitizenBodyResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Species     string            `json:"species"`
	Description string            `json:"description"`
	PhotoURL    string            `json:"photo_url"`
	Weight      float64           `json:"weight"`
	Height      float64           `json:"height"`
	HasPetHuman bool              `json:"has_pet_human"`
	Roles       RolesBodyResponse `json:"roles,omitempty"`
}

// RolesBodyResponse ...
type RolesBodyResponse []RoleBodyResponse

// RoleBodyResponse ...
type RoleBodyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
