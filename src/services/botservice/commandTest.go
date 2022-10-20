package botservice

import "fmt"

type CommandTest struct {
}

func (c *CommandTest) Execute() {
	fmt.Print("\n Test!")
}
