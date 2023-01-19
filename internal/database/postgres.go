package database

import (
	"BMinder/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func NewPostgres(cfg config.Database, logger zerolog.Logger) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SslMode,
	)

	db, err := sqlx.Open(cfg.Driver, dsn)
	if err != nil {
		logger.Error().Err(err).Msg("Ошибка с соединением с базой данных pgsql")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Error().Err(err).Msg("Ошибка проверки соединения с базой данных pgsql")
		return nil, err
	}

	return db, nil
}
