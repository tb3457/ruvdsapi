package cmd

import (
	"ruvdsapi/config"

	"github.com/Songmu/prompter"
)

var configureHandler = func(args []string) error {
	username := prompter.Prompt("Enter username", "user")
	password := prompter.Password("Enter password")
	apikey := prompter.Prompt("Enter apikey", "key")

	config := config.NewConfig()
	config.Set(username, password, apikey)
	return nil
}

func init() {
	rootCmd.AddCmd("configure", configureHandler)

}
