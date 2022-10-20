package observer

type Observer interface {
	Send(ChatID int64, message string)
}
