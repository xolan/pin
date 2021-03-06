package command

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// Command ...
// Struct representation of a pinned command
type Command struct {
	Command    string `json:"command"`
	Hash       string `json:"hash"`
	Identifier string `json:"identifier"`
}

// GetExecutable ...
// Get the executable portion of the command
func (command *Command) GetExecutable() string {
	var split = strings.Split(command.Command, " ")
	return split[0]
}

// GetArgs ...
// Get the args portion of the command
func (command *Command) GetArgs() string {
	var split = strings.Split(command.Command, " ")
	return strings.Join(split[1:], " ")
}

// Commands ...
// Struct representation of a list of pinned commands
type Commands struct {
	Commands []Command
}

// AddCommand - Add (append) a command to commands
func (commands *Commands) AddCommand(command *Command) *Commands {
	commands.Commands = append(commands.Commands, *command)
	return commands
}

// RemoveCommand - Remove a command from commands
func (commands *Commands) RemoveCommand(command *Command) *Commands {
	var remove []int
	for n, cmd := range commands.Commands {
		if cmd.Identifier == command.Identifier {
			remove = append(remove, n)
		}
	}
	for n := range remove {
		commands.Commands = append(commands.Commands[:n], commands.Commands[n+1:]...)
	}
	return commands
}

// GetCommand - Return the command with the given identifier
func (commands *Commands) GetCommand(identifier string) *Command {
	for _, cmd := range commands.Commands {
		if cmd.Identifier == identifier {
			return &cmd
		}
	}
	return nil
}

// HasCollision - Returns an error if there are any objects with the same identifier
func (commands *Commands) HasCollision() (bool, error) {
	var count map[string]int
	count = make(map[string]int)

	for _, command := range commands.Commands {
		count[command.Identifier]++
	}

	var err error
	for key, value := range count {
		if value > 1 {
			var errMsg = "Duplicate command entry found!"
			log.WithFields(log.Fields{
				"identifier": key,
				"count":      value,
			}).Error(errMsg)
			err = fmt.Errorf(errMsg)
		}
	}
	return err != nil, err
}
