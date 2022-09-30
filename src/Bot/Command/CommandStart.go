package Command

import "fmt"

type CommandStart struct {
	name string
}

func (c *CommandStart) execute() {
	fmt.Print("Execute")
}
