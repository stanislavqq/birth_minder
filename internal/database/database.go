package database

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/stanislavqq/birth_minder/internal/config"
)

func NewDatabase(cfg config.Database, logger zerolog.Logger) (*sqlx.DB, error) {
	switch cfg.Driver {
	case "postgres":
		return NewPostgres(cfg, logger)
	}

	err := errors.New("неизвестный драйвер")
	logger.Error().Err(err).Str("driver", cfg.Driver).Msg("Ошибка создания нового соединения с БД")

	return nil, err
}
