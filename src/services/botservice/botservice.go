package botservice

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type BotService struct {
	botapi tgbotapi.BotAPI
}

func NewBot(token string) (BotService, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return BotService{}, err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return BotService{botapi: *bot}, nil
}

func ListenCommands(bot BotService) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.botapi.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			cmdInputText := update.Message.Text
			cmdText := strings.TrimPrefix(cmdInputText, "/")

			factory := CommandFactory{}
			command, err := factory.CreateCommand(cmdText)
			if err != nil {
				fmt.Print(err)
			} else {
				command.Execute()
			}
		}
	}
}

func (b *BotService) SendTextToChat(ChatID int64, message string) {
	messageConf := tgbotapi.NewMessage(ChatID, message)
	b.botapi.Send(messageConf)
}
