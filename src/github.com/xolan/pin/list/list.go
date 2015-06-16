package list

import (
	"fmt"

	"github.com/xolan/pin/command"
	"github.com/xolan/pin/io"
)

var pinGlobal = new(command.Commands)
var pinLocal = new(command.Commands)

// Configuration flag set at runtime
var ShouldNotListGlobal bool

// Configuration flag set at runtime
var ShouldNotListLocal bool

// List pinned commands defined in ./.pin in found, else ~/.pin
func List() {
	if !ShouldNotListGlobal {
		io.ReadGlobal(pinGlobal)
	}

	if !ShouldNotListLocal {
		io.ReadLocal(pinLocal)
	}

	if !ShouldNotListGlobal || !ShouldNotListLocal {
		printPinned()
	}
}

func printPinned() {
	var hasGlobal = len(pinGlobal.Commands) > 0
	var hasLocal = len(pinLocal.Commands) > 0

	if hasGlobal {
		fmt.Println("Globally pinned commands")
		for _, command := range pinGlobal.Commands {
			fmt.Printf("[%s] %s: %s\n", command.Hash, command.Identifier, command.Command)
		}
	}

	if hasGlobal && hasLocal {
		fmt.Println("")
	}

	if hasLocal {
		fmt.Println("Locally pinned commands")
		for _, command := range pinLocal.Commands {
			fmt.Printf("[%s] %s: %s\n", command.Hash, command.Identifier, command.Command)
		}
	}
}
