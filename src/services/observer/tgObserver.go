package observer

import "BMinder/src/services/botservice"

type TGObserver struct {
	Bot botservice.BotService
}

func (o *TGObserver) Send(ChatID int64, message string) {
	o.Bot.SendTextToChat(ChatID, message)
}
