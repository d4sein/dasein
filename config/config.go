package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	// Config ..
	Config = getConfig("config/config.json")
)

// Config ...
type config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func getConfig(n string) (c config) {
	f, err := ioutil.ReadFile(n)
	json.Unmarshal(f, &c)

	if err != nil {
		log.Fatal(err)
	}

	return c
}
