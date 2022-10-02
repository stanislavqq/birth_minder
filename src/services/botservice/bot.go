package botservice

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Bot struct {
	botapi tgbotapi.BotAPI
}

func NewBot(token string) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return Bot{}, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return Bot{botapi: *bot}, nil
}

func (b *Bot) SendTextToChat(ChatID int64, message string) {
	messageConf := tgbotapi.NewMessage(ChatID, message)
	b.botapi.Send(messageConf)
}
