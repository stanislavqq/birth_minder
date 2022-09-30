package Command

type CommandFactory struct {
}

func (f *CommandFactory) Create(cmdName string) Command {
	var command Command

	if cmdName == "start" {
		command = &CommandStart{name: "start"}
	}

	return command
}
