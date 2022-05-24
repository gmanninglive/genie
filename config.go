package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config []Task

func LoadConfig(f Flags) Config {
  var config []Task
  var config_location string

  config_env, isEnvSet := os.LookupEnv("GENIE_CONFIG")

  __dir, err := os.UserHomeDir()
  Check(err)

  switch {
    case f.Config != "":
      config_location = strings.Replace(f.Config, "~", __dir, 1)
    case isEnvSet:
      config_location = strings.Replace(config_env, "~", __dir, 1)
    default:
      config_location = "config.json"
  }

  // Set global env
  GENIE.BASE = filepath.Dir(config_location)
  GENIE.Config = config_location

  fmt.Printf("Loading config from: %s\n", config_location)

  file, err := os.ReadFile(config_location)
  Check(err)

  json.Unmarshal([]byte(file), &config)

  return config
}
