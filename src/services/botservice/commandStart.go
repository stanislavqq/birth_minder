package botservice

import "fmt"

type CommandStart struct {
}

func (c *CommandStart) Execute() {
	fmt.Print("\nExecuted!")
}
