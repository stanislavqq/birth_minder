package telegram

import (
	"BMinder/internal/config"
	"BMinder/internal/notify"
	"github.com/rs/zerolog"
)

type TelegramNotifyProvider struct {
	bot    TelegramBot
	logger zerolog.Logger
	botCfg *config.TGBot
}

func New(config config.TGBot, logger zerolog.Logger) *TelegramNotifyProvider {
	bot, err := NewBot(config.Token)
	if err != nil {
		logger.Error().Err(err).Msg("Ошибка при создании бота")
	}
	return &TelegramNotifyProvider{bot: bot, logger: logger, botCfg: &config}
}

func (p *TelegramNotifyProvider) SendNotify(notify notify.Notify) (bool, error) {
	p.bot.SendTextToChat(int64(p.botCfg.NotifyChat), notify.Message)
	return true, nil
}
