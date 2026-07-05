package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"project-base-wahyu/config"
)

func NewPostgresConnection(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal konek ke database: %w", err)
	}

	
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database tidak merespon ping: %w", err)
	}

	return db, nil
}
