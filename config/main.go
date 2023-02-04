package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
)

var confDir = configdir.LocalConfig("ruvdscli")
var confFile = filepath.Join(confDir, "settings.json")

func checkPath() bool {
	var err error
	if _, err = os.Stat(confFile); os.IsNotExist(err) {
		return true
	}
	if err != nil {
		panic(err)
	}
	return true
}

type Config struct {
	ApiKey   string `json:"key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewConfig() Config {
	//first, lets check environment
	var cnf Config
	if username, ok := os.LookupEnv("RUVDS_USERNAME"); ok {
		if password, ok := os.LookupEnv("RUVDS_PASSWORD"); ok {
			if apikey, ok := os.LookupEnv("RUVDS_KEY"); ok {
				cnf.ApiKey = apikey
				cnf.Password = password
				cnf.Username = username
				return cnf
			}
		}
	}
	var err error
	var handler *os.File
	err = configdir.MakePath(confDir)
	if err != nil {
		panic(err)
	}
	//try to load first
	if handler, err = os.Open(confFile); os.IsNotExist(err) {
		return Config{}
	}
	if err != nil {
		panic(err)
	}
	defer handler.Close()
	decoder := json.NewDecoder(handler)
	decoder.Decode(&cnf)
	return cnf
}

func (c *Config) Set(username, password, apikey string) {
	changed := false
	if username != "" {
		changed = true
		c.Username = username
	}
	if password != "" {
		changed = true
		c.Password = password
	}
	if apikey != "" {
		changed = true
		c.ApiKey = apikey
	}
	if changed {
		c.sync()
	}
}

func (c *Config) sync() {
	handler, err := os.Create(confFile)
	if err != nil {
		panic(err)
	}
	defer handler.Close()

	encoder := json.NewEncoder(handler)
	encoder.Encode(&c)
}

func (c *Config) Show() string {
	data, _ := json.Marshal(c)
	return string(data)
}
