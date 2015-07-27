package do

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
	"github.com/xolan/pin/command"
	"github.com/xolan/pin/io"
)

// Configuration flag set at runtime
var IdentifierFlag string

var pinGlobal = new(command.Commands)
var pinLocal = new(command.Commands)

// Do - Execute pinned command if it exists
func Do() {
	io.ReadGlobal(pinGlobal)
	io.ReadLocal(pinLocal)

	var cmd = pinLocal.GetCommand(IdentifierFlag)
	if cmd == nil {
		log.Debugln("Not found locally")
	}
	cmd = pinGlobal.GetCommand(IdentifierFlag)
	if cmd == nil {
		log.Debugln("Not found globally")
	}

	if cmd == nil {
		log.Errorf("Command not found: %s", IdentifierFlag)
		os.Exit(1)
	} else {
		log.Infof("Command found: %s", IdentifierFlag)
		running := exec.Command(cmd.GetExecutable(), cmd.GetArgs())

		// Stdout buffer
		cmdOutput := &bytes.Buffer{}
		// Attach buffer to command
		running.Stdout = cmdOutput

		// Execute command
		err := running.Run() // will wait for command to return
		printError(err)
		// Only output the commands stdout
		printOutput(cmdOutput.Bytes())

		os.Exit(0)
	}

}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("%s\n", string(outs))
	}
}
