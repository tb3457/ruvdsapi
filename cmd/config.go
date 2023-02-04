package cmd

import (
	"fmt"
	"os"
)

var configCmd = NewCmd()

var configHandler = func(args []string) error {
	if len(args) > 0 {
		arg := args[0]
		if handler, ok := configCmd.cmd[arg]; ok {
			fmt.Fprintf(os.Stderr, "Looking for sub-handler")
			handler(args[1:])
		} else {
			fmt.Fprintf(os.Stderr, "Looking for action")
		}
	} else {
		fmt.Fprintf(os.Stderr, "aaaaa")
	}
	return nil
}

func init() {

	rootCmd.AddCmd("config", configHandler)
}
