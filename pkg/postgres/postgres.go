package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/marcelofabianov/picpay/config"
)

type Postgres struct {
	conn *pgx.Conn
}

func NewPostgres(conn *pgx.Conn) *Postgres {
	return &Postgres{conn: conn}
}

func (p *Postgres) Conn() *pgx.Conn {
	return p.conn
}

func (p *Postgres) Close(ctx context.Context) error {
	return p.conn.Close(ctx)
}

func (p *Postgres) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return p.conn.Ping(ctx)
}

func Connect(ctx context.Context, cfg config.DatabaseConfig) (*Postgres, error) {
	dns := formatDSN(cfg)

	conn, err := pgx.Connect(ctx, dns)
	if err != nil {
		return nil, err
	}

	return NewPostgres(conn), nil
}

func formatDSN(cfg config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SslMode)
}
