package command

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

// Command ...
// Struct representation of a pinned command
type Command struct {
	Command    string `json:"command"`
	Hash       string `json:"hash"`
	Identifier string `json:"identifier"`
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
			log.Errorf("Duplicate identifier found: \"%s\"", key)
			err = fmt.Errorf("Duplicate identifier found: \"%s\"", key)
		}
	}
	return err != nil, err
}
