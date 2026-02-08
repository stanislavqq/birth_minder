package telegram

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/stanislavqq/birth_minder/internal/config"
)

type CommandController struct {
	bot                TelegramBot
	logger             zerolog.Logger
	botCfg             *config.TGBot
	debug              bool
	commandHandlerList map[string]func(sender MessageSender)
}

func NewCommandController(config config.TGBot, debug bool, logger zerolog.Logger) *CommandController {
	bot, err := NewBot(config.Token, &logger)
	if err != nil {
		logger.Error().Err(err).Msg("Ошибка при создании бота")
	}
	return &CommandController{bot: bot, logger: logger, botCfg: &config, debug: debug, commandHandlerList: make(map[string]func(sender MessageSender))}
}

func (c *CommandController) HandleCommandFunc(command string, handlerFunc func(sender MessageSender)) {
	c.commandHandlerList[command] = handlerFunc
}

func (c *CommandController) Start() {
	for update := range c.bot.GetUpdates() {
		if update.Message != nil { // If we got a message

			c.logger.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				if update.Message.Command() == "status" {
					c.bot.SendTextToChat(update.Message.Chat.ID, "Bot is working.")
				}

				for command, handleFunc := range c.commandHandlerList {
					fmt.Println(command)
					if command == update.Message.Command() {
						handleFunc(&c.bot)
					}
				}

				return
			}
		}
	}
}
