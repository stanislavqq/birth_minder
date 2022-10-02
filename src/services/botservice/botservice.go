package botservice

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func ListenCommands(bot Bot) {
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
