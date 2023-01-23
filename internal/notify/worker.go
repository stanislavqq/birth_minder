package notify

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

type NotifyWorker struct {
	notifyCollector chan Notify
	msgProvider     NotifyProvider
	logger          zerolog.Logger
}

func NewWorker(notifyCollector chan Notify, provider NotifyProvider, logger zerolog.Logger) *NotifyWorker {
	return &NotifyWorker{
		notifyCollector: notifyCollector,
		msgProvider:     provider,
		logger:          logger,
	}
}

func (w *NotifyWorker) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		w.logger.Info().Msg("Notify worker is running")
		if err := w.listenNotifyMessages(); err != nil {
			w.logger.Error().Err(err).Msg("Failed running notify worker")
			w.stop()
			cancel()
		}
	}()

	select {
	case v := <-quit:
		w.logger.Info().Msgf("signal.Notify: %v", v)
		w.stop()
	case done := <-ctx.Done():
		w.logger.Info().Msgf("ctx.Done: %v", done)
		w.stop()
	}

	return nil
}

func (w *NotifyWorker) listenNotifyMessages() error {
	for {
		select {
		case notify := <-w.notifyCollector:
			w.logger.Debug().Fields(notify).Msg("Получили новое уведомление в работу")
			sended, err := w.msgProvider.SendNotify(notify)
			if err != nil {
				w.logger.Error().Err(err).Msg("Не удалось отправить уведомление")
				return err
			}
			if sended {
				w.logger.Debug().Msg("Уведомление отправлено")
			}
		}
	}
}

func (w *NotifyWorker) stop() {
	close(w.notifyCollector)
}
