package domain

type UserType string

type User struct {
	ID               ID               `json:"id" validate:"required,uuid"`
	Name             string           `json:"name" validate:"required,min=3,max=255,string"`
	Email            Email            `json:"email" validate:"required,email,max=150,string"`
	Password         Password         `json:"password" validate:"required,min=8,max=255,string"`
	DocumentRegistry DocumentRegistry `json:"document_registry" validate:"required,min=11,max=14,string"`
	Enabled          bool             `json:"enabled" validate:"required,bool"`
	CreatedAt        CreatedAt        `json:"created_at"`
	UpdatedAt        UpdatedAt        `json:"updated_at"`
	DeletedAt        *DeletedAt       `json:"deleted_at" validate:"omitempty"`
	Version          Version          `json:"version" validate:"required,numeric"`
}
