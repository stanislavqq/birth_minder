package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramBot struct {
	botapi tgbotapi.BotAPI
}

func NewBot(token string) (TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return TelegramBot{}, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return TelegramBot{botapi: *bot}, nil
}
func (b *TelegramBot) SendTextToChat(ChatID int64, message string) (tgbotapi.Message, error) {
	messageConf := tgbotapi.NewMessage(ChatID, message)
	send, err := b.botapi.Send(messageConf)
	if err != nil {
		return send, err
	}

	return send, nil
}
