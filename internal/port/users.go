package port

import "github.com/marcelofabianov/picpay/internal/domain"

// Errors

const (
	ErrUserAlreadyExists = "error_user_already_exists"
)

// PKG

type PasswordHasher interface {
	Hash(data string) (string, error)
	Compare(data, encodedHash string) (bool, error)
}

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

type UserService interface {
	CreateUserService
}

// Repository

type CreateUserRepositoryInput struct {
	User domain.User
}

type CreateUserRepositoryOutput struct {
	CreateUserRepositoryInput
}

type CreateUserRepository interface {
	CreateUser(input CreateUserRepositoryInput) (CreateUserRepositoryOutput, error)
	ExistsByEmailOrDocumentRegistry(email string, documentRegistry string) (bool, error)
}

type UserRepository interface {
	CreateUserRepository
}

// UseCase

type CreateUserInputUseCase struct {
	Name             string
	Email            string
	Password         string
	DocumentRegistry string
}

type CreateUserOutputUseCase struct {
	CreateUserRepositoryOutput
}

type CreateUserUseCase interface {
	Execute(input CreateUserInputUseCase) (CreateUserOutputUseCase, error)
}
