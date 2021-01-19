package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config ...
type Config struct {
	Token string `json:"token"`
}

func getConfig(n string) (config Config, err error) {
	f, err := ioutil.ReadFile(n)
	json.Unmarshal(f, &config)

	return config, err
}
