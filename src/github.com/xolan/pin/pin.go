package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xolan/pin/list"
	"os"
)

func sanity_check(v bool) {
	if v {
		fmt.Printf("Checking sanity...\n")
	}
	var path = os.ExpandEnv("$HOME/.pin")
	if _, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666); !os.IsNotExist(err) {
		fmt.Printf("Found ~/.pin\n")
		fmt.Errorf(err.Error())
	}
}

func main() {
	var Verbose bool

	var PinCmd = &cobra.Command{
		Use:   "pin",
		Short: "Pin is a command pinner, similar to aliasing",
		Long:  `Pin is a command pinner, similar to aliasing`,
		Run: func(cmd *cobra.Command, args []string) {
			sanity_check(Verbose)
		},
	}

	PinCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	var ListCmd = &cobra.Command{
		Use:   "list",
		Short: "Display a list of pinned commands",
		Long:  "Display a list of pinned commands",
		Run: func(cmd *cobra.Command, args []string) {
			list.List()
		},
	}
	PinCmd.AddCommand(ListCmd)
	PinCmd.Execute()
}
