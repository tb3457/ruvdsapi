package cmd

import (
	"fmt"
	"os"
)

type CmdHandler func([]string) error

type Cmd struct {
	cmd map[string]CmdHandler
}

func NewCmd() *Cmd {
	c := new(Cmd)
	c.cmd = make(map[string]CmdHandler)
	return c
}

var rootCmd = NewCmd()

func (c *Cmd) AddCmd(name string, handler CmdHandler) {
	c.cmd[name] = handler
}

func Exec() {
	cliArgs := os.Args[1:]
	if len(cliArgs) > 0 {
		if handler, ok := rootCmd.cmd[cliArgs[0]]; ok {
			handler(cliArgs[1:])
		} else {
			fmt.Fprintf(os.Stderr, "Command line argument not recognized\n")
		}
	}
}
