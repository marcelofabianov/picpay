package usecase

import (
	"context"
	"errors"

	"github.com/marcelofabianov/picpay/internal/domain"
	"github.com/marcelofabianov/picpay/internal/port"
)

type CreateUserUseCase struct {
	repository     port.CreateUserRepository
	passwordHasher port.PasswordHasher
}

func NewCreateUserUseCase(repository port.CreateUserRepository, passwordHasher port.PasswordHasher) port.CreateUserUseCase {
	return &CreateUserUseCase{
		repository:     repository,
		passwordHasher: passwordHasher,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input port.CreateUserInputUseCase) (port.CreateUserOutputUseCase, error) {
	exists, err := uc.userExists(ctx, string(input.Email), string(input.DocumentRegistry))
	if err != nil {
		return port.CreateUserOutputUseCase{}, err
	}
	if exists {
		return port.CreateUserOutputUseCase{}, errors.New(port.ErrUserAlreadyExists)
	}

	password, error := uc.passwordHasher.Hash(string(input.Password))
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
		ctx,
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

func (uc *CreateUserUseCase) userExists(ctx context.Context, email, documentRegistry string) (bool, error) {
	result, err := uc.repository.ExistsByEmailOrDocumentRegistry(ctx, email, documentRegistry)
	if err != nil {
		return false, err
	}

	if result {
		return true, nil
	}

	return false, nil
}
