package main

import (
	"BMinder/internal/config"
	"BMinder/internal/database"
	bevent "BMinder/internal/model/bvent"
	"BMinder/internal/notify"
	"BMinder/src/services/botservice"
	"BMinder/src/services/notifier"
	"BMinder/src/services/observer"
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

	if cfg.Project.Debug {
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

	bot, err := botservice.NewBot(cfg.TGBot.Token)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка инициализации бота")
	}

	notifyService := notifier.New()
	notifyService.SetObserver(&observer.TGObserver{Bot: bot})

	c := cron.New()
	rep := bevent.NewRepository(db, logger)
	job := notify.NewJob(rep, logger)
	job.Run()
	_, err = c.AddFunc("@every 5s", func() {
		job.Run()
	})
	if err != nil {
		panic(err)
	}
	defer c.Stop()
	c.Start()

	botservice.ListenCommands(bot)
}
