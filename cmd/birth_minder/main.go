package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stanislavqq/birth_minder/internal/api/server"
	"github.com/stanislavqq/birth_minder/internal/config"
	"github.com/stanislavqq/birth_minder/internal/database"
	"github.com/stanislavqq/birth_minder/internal/model/bevent"
	"github.com/stanislavqq/birth_minder/internal/notify"
	"github.com/stanislavqq/birth_minder/internal/personstore"
	"github.com/stanislavqq/birth_minder/internal/telegram"
)

const day = time.Hour * 24
const week = day * 7

func main() {
	migrations := flag.Bool("migration", false, "Define migrations start option")
	debug := flag.Bool("debug", false, "Define debug mode option")
	configFile := flag.String("file", "config.yml", "Set path config file")
	flag.Parse()

	if err := config.ReadConfigYML(*configFile); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	cfg := config.GetConfigInstance()
	logger := log.With().Logger()

	logger.Info().Msgf("Init app with = debug_mode: %s; cron: %s; message_format: %s", cfg.Debug, cfg.CronRule, cfg.FormatMessage)

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		os.Exit(0)
	}()

	if *debug {
		logger.Info().Msg("Debug mode on")
		cfg.Debug = true
	}

	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if _, err := os.Stat(cfg.Database.Sqlite.Path); err == nil {
		log.Error().Err(err).Msg("Файл бд существует")
	} else if errors.Is(err, os.ErrNotExist) {
		log.Error().Err(err).Msg("Файл бд не найден")
	}
	//db, dbErr := database.NewMysql(cfg.Database, logger)
	db, dbErr := database.NewSqlite(cfg.Database, logger)
	defer db.Close()

	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("Ошибка подключения к БД")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal().Err(dbErr).Msg("Ошибка ping к БД")
		return
	}

	if *migrations {
		if gooseError := goose.Up(db, cfg.Database.Migrations); gooseError != nil {
			log.Fatal().Err(gooseError).Msg("Ошибка выполнения миграции")
		}
		db.Close()
		return
	}

	notifyCollector := make(chan notify.Notify)

	c := cron.New()
	rep := bevent.NewRepository(db, logger)
	job := notify.NewJob(rep, cfg.FormatMessage, []time.Duration{day, week}, notifyCollector, cfg.Debug, logger)

	var cronRule string
	if len(cfg.CronRule) > 0 {
		cronRule = cfg.CronRule
	} else {
		cronRule = "0 10 * * * "
	}

	if cfg.Debug {
		job.Run()
		cronRule = "@every 1m"
	}

	logger.Info().Msg("Cron launched with rule: " + cronRule)

	_, err := c.AddFunc(cronRule, func() {
		job.Run()
	})
	if err != nil {
		panic(err)
	}
	defer c.Stop()
	c.Start()

	perStore := personstore.New(db, logger)

	if err := server.NewServer(perStore).Start(cfg, ctx, logger); err != nil {
		logger.Error().Err(err).Msg("Ошибка старта http сервера")
		return
	}

	provider := telegram.New(cfg.TGBot, cfg.Debug, logger)
	err = notify.NewWorker(notifyCollector, provider, logger).Start(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Не удалось запустить воркер")
	}

	cmdController := telegram.NewCommandController(cfg.TGBot, cfg.Debug, logger)
	cmdController.HandleCommandFunc("event", func(sender telegram.MessageSender) {
		t := time.Now()

		vday := t.Day()
		vmonth := int(t.Month())

		list, err := rep.GetUpcomingBDay(vday, vmonth)
		if err != nil {
			logger.Error().Err(err).Msg("Ошибка получения близжайших событий")
		}

		msg := "🤖 Ближайшие дни рождения в этом месяце: \n\n"
		if len(list) == 0 {
			msg = "🤖 Ближайшие дни рождения отсутствуют \n\n"
		}

		for _, v := range list {
			msg += v.GetFullName() + " - "
			msg += strconv.Itoa(int(v.Day)) + "." + strconv.Itoa(int(v.Month)) + "\n"
		}

		sender.SendTextToChat(int64(cfg.TGBot.NotifyChat), msg)
	})

	cmdController.HandleCommandFunc("stop", func(sender telegram.MessageSender) {
		cancel()
	})
	cmdController.Start()

	select {
	case v := <-quit:
		logger.Info().Msgf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		logger.Info().Msgf("ctx.Done: %v", done)
	}
}
