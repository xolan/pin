package list

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
)

type commands struct {
	Commands []struct {
		Command string `json:"command`
	}
}

var c = new(commands)

func List() {
	log.Debugln("listing...")
	var path, _ = homedir.Expand("~/.pin")
	if jsonFile, err := ioutil.ReadFile(path); err == nil {
		json.Unmarshal(jsonFile, &c)
	} else {
		log.Errorln(err.Error())
	}
	printPinned()
}

func printPinned() {
	for _, command := range c.Commands {
		fmt.Println(command.Command)
	}
}
