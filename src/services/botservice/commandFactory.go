package botservice

import (
	"errors"
)

type CommandFactory struct {
}

func (f *CommandFactory) CreateCommand(name string) (Command, error) {
	switch name {
	case "start":
		return &CommandStart{}, nil
	default:
		return nil, errors.New("command not exist")
	}
}
