package list

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
)

type commands struct {
	Commands []struct {
		Command string `json:"command"`
		Hash    string `json:"hash"`
	}
}

var c = new(commands)

func localFileExists() bool {
	var retval = false
	finfo, err := os.Stat("./.pin")
	if err != nil {
		log.Debugln("no local .pin")
		retval = false
		return retval
	}
	if finfo.IsDir() {
		log.Debugln("no local .pin")
		retval = false
	} else {
		log.Debugln("local .pin")
		retval = true
	}
	return retval
}

// List pinned commands defined in ./.pin in found, else ~/.pin
func List() {
	log.Debugln("listing from home...")
	var path, _ = homedir.Expand("~/.pin")
	if jsonFile, err := ioutil.ReadFile(path); err == nil {
		json.Unmarshal(jsonFile, &c)
	} else {
		log.Errorln(err.Error())
	}
	log.Debugln("listing local...")
	if localFileExists() {
		if jsonFile, err := ioutil.ReadFile("./.pin"); err == nil {
			json.Unmarshal(jsonFile, &c)
		} else {
			log.Errorln(err.Error())
		}
	}

	printPinned()
}

func printPinned() {
	for n, command := range c.Commands {
		var hash = sha1.Sum([]byte(command.Command))
		fmt.Printf("[%x] %d: %s\n", hash, n, command.Command)
	}
}
