package Bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	ID    int
	Token string
	bot   *tgbotapi.BotAPI
}

type Config struct {
	Token string
	Debug bool
}

func Create(c Config) (Bot, error) {
	bot, err := tgbotapi.NewBotAPI(c.Token)
	bot.Debug = c.Debug

	if err != nil {
		return nil, err
	}

	Bot := Bot{
		ID:    0,
		Token: "",
		bot:   nil,
	}

	return Bot, nil
}

func (b *Bot) ListenCommands() {
	bot := b.bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			//fmt.Println(update.Message.IsCommand(), update.Message.Text)

			if update.Message.Text == "/start" {
				stmt, err := db.Prepare("INSERT INTO accounts(id) VALUES(?)")

				checkErr(err)
				fmt.Println(update.Message.From.ID)
				res, err := stmt.Exec(update.Message.From.ID)
				checkErr(err)

				id, err := res.LastInsertId()
				checkErr(err)

				fmt.Println(id)
				message := tgbotapi.NewMessage(update.Message.Chat.ID, "Success register!")
				bot.Send(message)
			}

		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
