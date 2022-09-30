package Command

import (
	"BMinder/src/Bot/Command"
	"reflect"
	"testing"
)

func TestCommandFactory(t *testing.T) {
	factory := Command.CommandFactory{}

	t.Log("Given the need to create command with name")
	{
		testID := 1
		t.Logf("\tTest %d:\tWhen working crating command and name is start", testID)
		{
			command := factory.Create("start")
			startCommand := Command.CommandStart{}
			if reflect.TypeOf(command) != reflect.TypeOf(startCommand) {
				t.Errorf("CommandFactory.Create() must return implementing Command interface. %s given.", command)
				t.Error(startCommand)
			}

			//t.Log("Ok")
		}
	}
}
