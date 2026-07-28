package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/rs/zerolog"
	"github.com/stanislavqq/birth_minder/internal/config"
)

func NewSqlite(database config.Database, logger zerolog.Logger) (*sql.DB, error) {
	cfg := database.Sqlite
	fmt.Println(cfg)
	db, err := sql.Open("sqlite3", "./db_bminder.sqlite")
	if err != nil {
		logger.Error().Str("err", err.Error()).Msg("Неудалось подключиться к sqlite бд")
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
