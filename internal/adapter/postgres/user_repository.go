package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/marcelofabianov/picpay/internal/domain"
	"github.com/marcelofabianov/picpay/internal/port"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) port.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, input port.CreateUserRepositoryInput) (port.CreateUserRepositoryOutput, error) {
	query := `
		INSERT INTO users (id, name, email, password, document_registry, enabled, created_at, updated_at, version)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, name, email, document_registry, created_at, updated_at, enabled
	`

	var user domain.User

	err := r.db.QueryRow(
		ctx,
		query,
		input.User.ID,
		input.User.Name,
		input.User.Email,
		input.User.Password,
		input.User.DocumentRegistry,
		input.User.Enabled,
		input.User.CreatedAt,
		input.User.UpdatedAt,
		input.User.Version,
	).Scan(&user.ID, &user.Name, &user.Email, &user.DocumentRegistry, &user.CreatedAt, &user.UpdatedAt, &user.Enabled)

	if err != nil {
		return port.CreateUserRepositoryOutput{}, err
	}

	return port.CreateUserRepositoryOutput{
		CreateUserRepositoryInput: port.CreateUserRepositoryInput{
			User: user,
		},
	}, nil
}

func (r *UserRepository) ExistsByEmailOrDocumentRegistry(ctx context.Context, email string, documentRegistry string) (bool, error) {
	query := `
		SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 OR document_registry = $2)
	`

	var exists bool

	err := r.db.QueryRow(ctx, query, email, documentRegistry).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
