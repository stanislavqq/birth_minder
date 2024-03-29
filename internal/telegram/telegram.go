package telegram

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/stanislavqq/birth_minder/internal/config"
	"github.com/stanislavqq/birth_minder/internal/notify"
)

type TelegramNotifyProvider struct {
	bot    TelegramBot
	logger zerolog.Logger
	botCfg *config.TGBot
	debug  bool
}

func New(config config.TGBot, debug bool, logger zerolog.Logger) *TelegramNotifyProvider {
	bot, err := NewBot(config.Token)
	if err != nil {
		logger.Error().Err(err).Msg("Ошибка при создании бота")
	}
	return &TelegramNotifyProvider{bot: bot, logger: logger, botCfg: &config, debug: debug}
}

func (p *TelegramNotifyProvider) SendNotify(notify notify.Notify) (bool, error) {
	chatId := int64(p.botCfg.NotifyChat)
	message, err := p.bot.SendTextToChat(chatId, notify.Message)
	if err != nil {
		p.logger.Error().Err(err).Str("message", notify.Message).Int64("chat_id", chatId).Msg("Не удалось отправить уведомление")
		return false, err
	}

	if p.debug {
		fmt.Println(message)
		p.logger.Debug().Int("chat_id", p.botCfg.NotifyChat).Msg("Сообщение отправлено")
	}

	return true, nil
}
