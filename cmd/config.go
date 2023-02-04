package cmd

import (
	"fmt"
	"os"
)

var configCmd = NewCmd()

func usage() {
	fmt.Fprintf(os.Stderr, "Your command not found. But you can use this sub-cmd:\n")
	for name := range configCmd.cmd {
		fmt.Fprintf(os.Stderr, "%s\n", name)
	}
	os.Exit(1)
}

var configHandler = func(args []string) error {
	if len(args) > 0 {
		arg := args[0]
		if handler, ok := configCmd.cmd[arg]; ok {
			handler(args[1:])
		} else {
			usage()
		}
	} else {
		usage()
	}
	return nil
}

func init() {

	rootCmd.AddCmd("config", configHandler)
}
