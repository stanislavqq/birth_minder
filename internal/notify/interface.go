package notify

type NotifyProvider interface {
	SendNotify(notify Notify) (bool, error)
}

type Job interface {
	Run()
}
