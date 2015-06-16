package command

import "testing"

func TestCollisions(t *testing.T) {
	var cmd1 = &Command{
		Identifier: "a",
		Command:    "cmd1",
	}

	var cmd2 = &Command{
		Identifier: "a",
		Command:    "cmd2",
	}

	var cmd3 = &Command{
		Identifier: "b",
		Command:    "cmd3",
	}

	var cmds = new(Commands)
	cmds.AddCommand(cmd1)
	cmds.AddCommand(cmd2)

	if _, err := cmds.HasCollision(); err != nil {
		t.Log("Got expected collision")
	}

	cmds = new(Commands)
	cmds.AddCommand(cmd1)
	cmds.AddCommand(cmd3)

	if _, err := cmds.HasCollision(); err != nil {
		t.Error(err)
	}
}
