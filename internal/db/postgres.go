package db

import (
	"context"
	"otp-auth/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres(url string) (*Postgres, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return &Postgres{Pool: pool}, nil
}

func (p *Postgres) FindOrCreateUser(ctx context.Context, identifier string) (*models.User, error) {
	var u models.User
	err := p.Pool.QueryRow(ctx,
		`INSERT INTO users(identifier, verified, created_at, last_login)
         VALUES ($1, false, NOW(), NOW())
         ON CONFLICT (identifier) DO UPDATE SET last_login=NOW()
         RETURNING id, identifier, verified, created_at, last_login`,
		identifier).Scan(&u.ID, &u.Identifier, &u.Verified, &u.CreatedAt, &u.LastLogin)
	return &u, err
}

func (p *Postgres) UpdateVerified(ctx context.Context, id string) error {
	_, err := p.Pool.Exec(ctx, `UPDATE users SET verified=true, last_login=NOW() WHERE id=$1`, id)
	return err
}
