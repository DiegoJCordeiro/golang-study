package dto

import "github.com/google/uuid"

type RoleDto struct {
	ID          uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func NewRoleDto(id uuid.UUID, name string, description string) *RoleDto {

	return &RoleDto{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
