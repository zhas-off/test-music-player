package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

type PlaylistDatabase struct {
	Conn   *pgx.Conn
	Logger *zerolog.Logger
}

// Инициализация БД
func InitDatabase(url string) (*PlaylistDatabase, error) {
	var db PlaylistDatabase

	dsn := os.Getenv("POSTGRES_URI")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil || dsn == "" {
		conn, err = pgx.Connect(context.Background(), url)
		if err != nil {
			return &db, err
		}
	}
	_, err = conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS playlist (
		song VARCHAR(1000) NOT NULL,
		duration VARCHAR(50) NOT NULL
	);`)
	if err != nil {
		return nil, fmt.Errorf("table creating error: %v", err)
	}

	db.Conn = conn
	return &db, nil
}
