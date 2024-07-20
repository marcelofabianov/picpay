package domain

type UserType string

type User struct {
	ID               ID
	Name             string
	Email            Email
	Password         Password
	DocumentRegistry DocumentRegistry
	Enabled          bool
	CreatedAt        CreatedAt
	UpdatedAt        UpdatedAt
	DeletedAt        *DeletedAt
	Version          Version
}
