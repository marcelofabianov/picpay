package port

import "github.com/marcelofabianov/picpay/internal/domain"

// Request

type UserCreateRequest struct {
	Name             string `json:"name" validate:"required,min=3,max=100"`
	Email            string `json:"email" validate:"required,email,min=3,max=150"`
	Password         string `json:"password" validate:"required,min=6,max=100"`
	DocumentRegistry string `json:"document_registry" validate:"required,min=11,max=14"`
}

// Presenter

type UserPresenter struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	DocumentRegistry string `json:"document_registry"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	Enabled          bool   `json:"enabled"`
}

// Service

type CreateUserService interface {
	CreateUser(request UserCreateRequest) (UserPresenter, error)
}

// UseCase

type CreateUserInputUseCase struct {
	Name             string
	Email            string
	Password         string
	DocumentRegistry string
}

type CreateUserOutputUseCase struct {
	User domain.User
}

type CreateUserUseCase interface {
	Execute(input CreateUserInputUseCase) (CreateUserOutputUseCase, error)
}
