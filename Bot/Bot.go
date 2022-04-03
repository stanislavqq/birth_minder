package Bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Create(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI("311927713:AAE8BOsoajS7TTMU87swuEfkPhmIlBV5_Xo")
	if err != nil {
		log.Panic(err)
	}

	return bot
}
