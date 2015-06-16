package add

import (
	"crypto/sha1"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/xolan/pin/command"
	"github.com/xolan/pin/io"
)

// Configuration flag set at runtime
var CommandFlag string

// Configuration flag set at runtime
var IdentifierFlag string

// Configuration flag set at runtime
var LocalFlag bool

func validateFlags() bool {
	if CommandFlag == "" {
		return false
	}
	if IdentifierFlag == "" {
		return false
	}
	return true
}

// Add command
func Add() {
	if !validateFlags() {
		log.Errorln("At least cmd and identifier flags must be set")
		os.Exit(1)
	}

	log.Debugln("Adding command")
	log.Debugf("    cmd: %s", CommandFlag)
	log.Debugf("     id: %s", IdentifierFlag)
	log.Debugf("  local: %s", LocalFlag)

	var hash = sha1.Sum([]byte(CommandFlag))

	cmd := &command.Command{
		Command:    CommandFlag,
		Identifier: IdentifierFlag,
		Hash:       fmt.Sprintf("%x", hash),
	}

	io.Store(cmd, LocalFlag)
}
