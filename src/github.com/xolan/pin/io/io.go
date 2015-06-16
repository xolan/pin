package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
	"github.com/xolan/pin/command"
)

// LocalFileExists - Return true if ./.pin exists
func LocalFileExists() bool {
	var retval = false
	finfo, err := os.Stat("./.pin")
	if err != nil {
		log.Debugln("local .pin not found")
		return false
	}
	if finfo.IsDir() {
		log.Debugln("local .pin not found")
		retval = false
	} else {
		log.Debugln("local .pin found")
		retval = true
	}
	return retval
}

// ReadGlobal - Reads ~/.pin into a *command.Commands
func ReadGlobal(into *command.Commands) {
	log.Debugln("reading from ~/.pin")
	var path, _ = homedir.Expand("~/.pin")
	if jsonFile, err := ioutil.ReadFile(path); err == nil {
		json.Unmarshal(jsonFile, &into)
	} else {
		log.Errorln(err.Error())
	}
}

// ReadLocal - Reads ./.pin into a *command.Commands
func ReadLocal(into *command.Commands) {
	log.Debugln("reading from ./.pin")
	if LocalFileExists() {
		if jsonFile, err := ioutil.ReadFile("./.pin"); err == nil {
			json.Unmarshal(jsonFile, &into)
		} else {
			log.Errorln(err.Error())
		}
	}
}

// Store cmd into a .pin-file, either globally (~/.pin) for locally (./.pin)
func Store(cmd *command.Command, locally bool) {
	var stored = new(command.Commands)
	var path string

	if locally {
		ReadLocal(stored)
		path = "./.pin"
	} else {
		ReadGlobal(stored)
		path, _ = homedir.Expand("~/.pin")
	}

	stored.AddCommand(cmd)

	if _, err := stored.HasCollision(); err != nil {
		fmt.Errorf("Collision found")
	}

	marshalled, _ := json.MarshalIndent(stored, "", "  ")
	log.Debugln(string(marshalled))

	if err := ioutil.WriteFile(path, marshalled, 0644); err == nil {
		log.Debugln("Successfully written to file")
	} else {
		log.Errorln(err.Error())
	}
}
