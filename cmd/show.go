package cmd

import (
	"fmt"
	"ruvdsapi/config"
)

var showHandler = func(args []string) error {
	config := config.NewConfig()
	fmt.Println(config.Show())
	return nil
}

func init() {
	configCmd.AddCmd("show", showHandler)
}
