package main

import (
	"BMinder/internal/config"
	"BMinder/internal/database"
	bevent "BMinder/internal/model/bevent"
	"BMinder/internal/notify"
	"BMinder/internal/telegram"
	"flag"
	"github.com/pressly/goose/v3"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

const day = time.Hour * 24
const week = day * 7

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	migrations := flag.Bool("migration", false, "Define migrations start option")
	debug := flag.Bool("debug", false, "Define debug mode option")
	flag.Parse()

	cfg := config.GetConfigInstance()
	logger := log.With().Logger()

	if *debug {
		cfg.Debug = true
	}

	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	db, dbErr := database.NewDatabase(cfg.Database, logger)
	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("Ошибка подключения к БД")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal().Err(dbErr).Msg("Ошибка ping к БД")
		return
	}

	if *migrations {
		if gooseError := goose.Up(db.DB, cfg.Database.Migrations); gooseError != nil {
			log.Fatal().Err(gooseError).Msg("Ошибка выполнения миграции")
		}
		return
	}

	notifyCollector := make(chan notify.Notify)

	c := cron.New()
	rep := bevent.NewRepository(db, logger)
	job := notify.NewJob(rep, cfg.FormatMessage, []time.Duration{day, week}, notifyCollector, cfg.Debug, logger)

	cronRule := "@daily"

	if cfg.Debug {
		job.Run()
		cronRule = "@every 10s"
	}

	_, err := c.AddFunc(cronRule, func() {
		job.Run()
	})
	if err != nil {
		panic(err)
	}
	defer c.Stop()
	c.Start()

	provider := telegram.New(cfg.TGBot, logger)
	err = notify.NewWorker(notifyCollector, provider, logger).Start()
	if err != nil {
		logger.Error().Err(err).Msg("Не удалось запустить воркер")
	}
}
