package usecase

import (
	"errors"

	"github.com/marcelofabianov/picpay/internal/domain"
	"github.com/marcelofabianov/picpay/internal/port"
)

type CreateUserUseCase struct {
	repository     port.CreateUserRepository
	passwordHasher port.PasswordHasher
}

func NewCreateUserUseCase(repository port.CreateUserRepository) port.CreateUserUseCase {
	return &CreateUserUseCase{
		repository: repository,
	}
}

func (uc *CreateUserUseCase) Execute(input port.CreateUserInputUseCase) (port.CreateUserOutputUseCase, error) {
	exists, err := uc.UserExists(input.Email, input.DocumentRegistry)
	if err != nil {
		return port.CreateUserOutputUseCase{}, err
	}
	if exists {
		return port.CreateUserOutputUseCase{}, errors.New(port.ErrUserAlreadyExists)
	}

	password, error := uc.passwordHasher.Hash(input.Password)
	if error != nil {
		return port.CreateUserOutputUseCase{}, error
	}

	user := domain.User{
		ID:               domain.NewID(),
		Name:             input.Name,
		Email:            domain.Email(input.Email),
		Password:         domain.Password(password),
		DocumentRegistry: domain.DocumentRegistry(input.DocumentRegistry),
		Enabled:          true,
		CreatedAt:        domain.CreatedAtNow(),
		UpdatedAt:        domain.UpdatedAtNow(),
		Version:          domain.InitVersion(),
	}

	output, err := uc.repository.CreateUser(
		port.CreateUserRepositoryInput{
			User: user,
		},
	)

	if err != nil {
		return port.CreateUserOutputUseCase{}, err
	}

	return port.CreateUserOutputUseCase{
		CreateUserRepositoryOutput: output,
	}, nil
}

func (uc *CreateUserUseCase) UserExists(email, documentRegistry string) (bool, error) {
	result, err := uc.repository.ExistsByEmailOrDocumentRegistry(email, documentRegistry)
	if err != nil {
		return false, err
	}

	if result {
		return true, nil
	}

	return false, nil
}
