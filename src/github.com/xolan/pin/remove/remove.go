package remove

import (
	log "github.com/Sirupsen/logrus"
	"github.com/xolan/pin/command"
	"github.com/xolan/pin/io"
)

// Configuration flag set at runtime
var IdentifierFlag string

// Configuration flag set at runtime
var LocalFlag bool

// Remove - Removes (and saves) a command with the supplied identifier
func Remove(identifier *string, local *bool) {
	var stored = new(command.Commands)
	if *local {
		io.ReadLocal(stored)
	} else {
		io.ReadGlobal(stored)
	}

	var cmd = stored.GetCommand(*identifier)
	if cmd != nil {
		log.Debugln("Removing command", cmd)
		stored.RemoveCommand(cmd)
		io.Store(stored, *local)
	}
}
