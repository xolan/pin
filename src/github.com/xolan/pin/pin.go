package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/xolan/pin/list"
)

func sanity_check() {
	log.Debugln("Checking sanity...")
	var path, _ = homedir.Expand("~/.pin")
	if _, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666); !os.IsNotExist(err) {
		log.Debugf("Found %s", path)
	} else {
		log.Errorln(err.Error())
	}
}

func config(v bool) {
	if v {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}
	sanity_check()
}

func main() {
	var Verbose bool

	var PinCmd = &cobra.Command{
		Use:   "pin",
		Short: "Pin is a command pinner, similar to aliasing",
		Long:  `Pin is a command pinner, similar to aliasing`,
		Run: func(cmd *cobra.Command, args []string) {
			config(Verbose)
			fmt.Println(cmd.UsageString())
		},
	}

	var ListCmd = &cobra.Command{
		Use:   "list",
		Short: "Display a list of pinned commands",
		Long:  "Display a list of pinned commands",
		Run: func(cmd *cobra.Command, args []string) {
			config(Verbose)
			list.List()
		},
	}

	var GenDocsCmd = &cobra.Command{
		Use:   "gendocs",
		Short: "Generate documentation for this program",
		Long:  "Generate documentation for this program",
		Run: func(cmd *cobra.Command, args []string) {
			config(Verbose)
			cobra.GenMarkdownTree(PinCmd, "./")
		},
	}

	PinCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	PinCmd.AddCommand(ListCmd)
	PinCmd.AddCommand(GenDocsCmd)
	PinCmd.Execute()
}
