package application

import (
	"context"

	"github.com/marcelofabianov/picpay/internal/domain"
	"github.com/marcelofabianov/picpay/internal/port"
)

type UserService struct {
	createUser port.CreateUserUseCase
}

func NewUserService(createUser port.CreateUserUseCase) port.UserService {
	return &UserService{
		createUser: createUser,
	}
}

func (s *UserService) CreateUser(ctx context.Context, request port.UserCreateRequest) (port.UserPresenter, error) {
	useCaseOutput, err := s.createUser.Execute(
		ctx,
		port.CreateUserInputUseCase{
			Name:             request.Name,
			Email:            domain.Email(request.Email),
			Password:         domain.Password(request.Password),
			DocumentRegistry: domain.DocumentRegistry(request.DocumentRegistry),
		},
	)

	if err != nil {
		return port.UserPresenter{}, err
	}

	user := useCaseOutput.User

	return port.UserPresenter{
		ID:               string(user.ID),
		Name:             user.Name,
		Email:            string(user.Email),
		DocumentRegistry: string(user.DocumentRegistry),
		CreatedAt:        user.CreatedAt.Format(),
		UpdatedAt:        user.CreatedAt.Format(),
		Enabled:          user.Enabled,
	}, nil
}
