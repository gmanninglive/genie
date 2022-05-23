package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config []Task

func LoadConfig(config_location string) Config {
  var config []Task
  GENIE.BASE = filepath.Dir(config_location)

  f, err := os.ReadFile(config_location)
  Check(err)

  json.Unmarshal([]byte(f), &config)

  return config
}
