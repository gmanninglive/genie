package main

import (
	"encoding/json"
	"os"
)

type Config []Task

func LoadConfig(config_location string) Config {
  var config []Task

  f, err := os.ReadFile(config_location)
  Check(err)

  json.Unmarshal([]byte(f), &config)

  return config
}
