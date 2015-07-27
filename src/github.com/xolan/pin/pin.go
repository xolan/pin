package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/xolan/pin/add"
	"github.com/xolan/pin/do"
	"github.com/xolan/pin/list"
	"github.com/xolan/pin/remove"
)

func sanityCheck() {
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
	sanityCheck()
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

	var DoCmd = &cobra.Command{
		Use:   "do",
		Short: "Execute command with the given identifier",
		Long:  "Execute command with the given identifier",
		Run: func(cmd *cobra.Command, args []string) {
			config(Verbose)
			do.Do()
		},
	}

	DoCmd.Flags().StringVarP(&do.IdentifierFlag, "identifier", "i", "", "Which command to execute")

	var ListCmd = &cobra.Command{
		Use:   "list",
		Short: "Display a list of pinned commands",
		Long:  "Display a list of pinned commands",
		Run: func(cmd *cobra.Command, args []string) {
			config(Verbose)
			list.List()
		},
	}

	ListCmd.Flags().BoolVarP(&list.ShouldNotListGlobal, "no-global", "", false, "Do not read pinned commands from ~/.pin")
	ListCmd.Flags().BoolVarP(&list.ShouldNotListLocal, "no-local", "", false, "Do not read pinned commands from ./.pin")

	var AddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add (pin) a command",
		Long:  "Add (pin) a command",
		Run: func(cmd *cobra.Command, args []string) {
			config(Verbose)
			add.Add()
		},
	}

	AddCmd.Flags().StringVarP(&add.CommandFlag, "cmd", "c", "", "What command to pin")
	AddCmd.Flags().StringVarP(&add.IdentifierFlag, "identifier", "i", "", "What to pin the command as")
	AddCmd.Flags().BoolVarP(&add.LocalFlag, "local", "l", false, "Whether to store the pinned command locally")

	var RemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove (unpin) a command",
		Long:  "Remove (unpin) a command",
		Run: func(cmd *cobra.Command, args []string) {
			config(Verbose)
			remove.Remove(&remove.IdentifierFlag, &remove.LocalFlag)
		},
	}

	RemoveCmd.Flags().StringVarP(&remove.IdentifierFlag, "identifier", "i", "", "Which command to remove")
	RemoveCmd.Flags().BoolVarP(&remove.LocalFlag, "local", "l", false, "Whether to remove a locally pinned command")

	var GenDocsCmd = &cobra.Command{
		Use:   "gendocs",
		Short: "Generate documentation for this program",
		Long:  "Generate documentation for this program",
		Run: func(cmd *cobra.Command, args []string) {
			config(Verbose)
			cobra.GenMarkdownTree(PinCmd, "./docs")
		},
	}

	PinCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	PinCmd.AddCommand(DoCmd)
	PinCmd.AddCommand(AddCmd)
	PinCmd.AddCommand(ListCmd)
	PinCmd.AddCommand(RemoveCmd)
	PinCmd.AddCommand(GenDocsCmd)
	PinCmd.Execute()
}
