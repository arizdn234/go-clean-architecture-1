package repository

import (
	"context"

	"github.com/arizdn234/go-clean-architecture-1/internal/domain"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Save(user domain.User) error
}

type userRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (r *userRepository) Save(user domain.User) error {
	query := `
		INSERT INTO users(name)
		VALUES($1)
	`

	_, err := r.conn.Exec(
		context.Background(),
		query,
		user.Name,
	)

	return err
}