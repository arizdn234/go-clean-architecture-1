package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/arizdn234/go-clean-architecture-1/internal/config"
)

func NewPostgres(cfg config.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	conn, err := pgx.Connect(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	return conn, nil
}