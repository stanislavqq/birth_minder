package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
)

type MessageSender interface {
	SendTextToChat(ChatID int64, message string) (tgbotapi.Message, error)
}

type TelegramBot struct {
	botapi tgbotapi.BotAPI
	logger *zerolog.Logger
}

func NewBot(token string, logger *zerolog.Logger) (TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return TelegramBot{}, err
	}

	return TelegramBot{botapi: *bot, logger: logger}, nil
}
func (b *TelegramBot) SendTextToChat(ChatID int64, message string) (tgbotapi.Message, error) {
	messageConf := tgbotapi.NewMessage(ChatID, message)
	send, err := b.botapi.Send(messageConf)
	if err != nil {
		return send, err
	}

	return send, nil
}

func (b *TelegramBot) GetUpdates() tgbotapi.UpdatesChannel {
	bot := b.botapi

	bot.Debug = true

	b.logger.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	return updates
	//for update := range updates {
	//	if update.Message != nil { // If we got a message
	//
	//		b.logger.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//		if update.Message.IsCommand() {
	//			if update.Message.Command() == "status" {
	//				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bot is working. ")
	//				//msg.ReplyToMessageID = update.Message.MessageID
	//
	//				bot.Send(msg)
	//			}
	//
	//			if update.Message.Command() == "event" {
	//				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Bot is working. ")
	//				//msg.ReplyToMessageID = update.Message.MessageID
	//
	//				bot.Send(msg)
	//			}
	//
	//			return
	//		}
	//
	//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//		msg.ReplyToMessageID = update.Message.MessageID
	//
	//		bot.Send(msg)
	//	}
	//}
}
