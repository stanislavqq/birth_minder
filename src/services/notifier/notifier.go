package notifier

import (
	"BMinder/src/services/observer"
	"fmt"
)

type Notifier interface {
	SetObserver(observer observer.Observer)
	Notify()
}

type Observers []observer.Observer

type BirthdayNotifier struct {
	items Observers
}

func New() BirthdayNotifier {
	return BirthdayNotifier{
		items: Observers{},
	}
}

func (n *BirthdayNotifier) Notify(ChatID int64, message string) {
	for _, observerItem := range n.items {
		observerItem.Send(ChatID, message)
		fmt.Print("\n Send")
	}
}

func (n *BirthdayNotifier) SetObserver(observer observer.Observer) {
	n.items = append(n.items, observer)
	fmt.Print("Added observer", n.items)
}
