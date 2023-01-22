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
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	migrations := flag.Bool("migration", false, "Define migrations start option")
	flag.Parse()

	cfg := config.GetConfigInstance()
	logger := log.With().Logger()

	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	db, dbErr := database.NewDatabase(cfg.Database, logger)

	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("Ошибка подключения к БД")
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
	job := notify.NewJob(rep, notifyCollector, logger)

	_, err := c.AddFunc("@every 5s", func() {
		job.Run()
	})
	if err != nil {
		panic(err)
	}

	provider := telegram.New(cfg.TGBot, logger)
	worker := notify.NewWorker(notifyCollector, provider, logger)

	defer func() {
		c.Stop()
		worker.Stop()
	}()

	c.Start()
	worker.Start()
}
