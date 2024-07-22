package internal

import (
	"github.com/jackc/pgx/v5"
	"github.com/marcelofabianov/picpay/internal/adapter/postgres"
	"github.com/marcelofabianov/picpay/internal/application"
	"github.com/marcelofabianov/picpay/internal/domain/usecase"
	"github.com/marcelofabianov/picpay/internal/port"
	"github.com/marcelofabianov/picpay/pkg/hasher"
	"go.uber.org/dig"
)

type Container struct {
	*dig.Container
}

func NewContainer(db *pgx.Conn) *Container {
	container := dig.New()

	registerPackages(container)
	registerRepository(container, db)
	registerUseCase(container)
	registerService(container)

	return &Container{Container: container}
}

func registerPackages(container *dig.Container) {
	container.Provide(func() port.PasswordHasher {
		return hasher.NewHasher()
	})
}

func registerRepository(container *dig.Container, db *pgx.Conn) {
	container.Provide(func() port.UserRepository {
		return postgres.NewUserRepository(db)
	})
}

func registerUseCase(container *dig.Container) {
	container.Provide(func(repository port.UserRepository, hasher port.PasswordHasher) port.CreateUserUseCase {
		return usecase.NewCreateUserUseCase(repository, hasher)
	})
}

func registerService(container *dig.Container) {
	container.Provide(func(createUser port.CreateUserUseCase) port.UserService {
		return application.NewUserService(createUser)
	})
}
