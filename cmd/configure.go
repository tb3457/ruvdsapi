package cmd

import (
	"os"
	"ruvdsapi/config"

	"github.com/Songmu/prompter"
)

var configureHandler = func(args []string) error {
	var username, password, apikey string
	var ok bool
	if username, ok = os.LookupEnv("RUVDS_USERNAME"); !ok {
		if password, ok = os.LookupEnv("RUVDS_PASSWORD"); !ok {
			if apikey, ok = os.LookupEnv("RUVDS_KEY"); !ok {
				username = prompter.Prompt("Enter username", "user")
				password = prompter.Password("Enter password")
				apikey = prompter.Prompt("Enter apikey", "key")
			}
		}
	}
	config := config.Config{}
	config.Set(username, password, apikey)
	return nil
}

func init() {
	rootCmd.AddCmd("configure", configureHandler)

}
