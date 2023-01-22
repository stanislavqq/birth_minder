package notify

import (
	"time"
)

type Notify struct {
	AfterWhichTime time.Duration
	Message        string
}

//func (n *Notify) CreateMessage(message string) string {
//	n.Message = message
//}

func NewNotify(message string, dur time.Duration) Notify {
	return Notify{
		AfterWhichTime: dur,
		Message:        message,
	}
}
