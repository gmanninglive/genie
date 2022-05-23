package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Env struct {
  Config string
  BASE string
}

var GENIE Env

func Check(e error) {
  if e != nil {
      panic(e)
  }
}

func readflags() {
  var config_location string
  config_env, isEnvSet := os.LookupEnv("GENIE_CONFIG")

  __dir, err := os.UserHomeDir()
  Check(err)

  configPtr := flag.String("config", "", "set location of config.json:\n - Include filename,\n - Custom filnames are accepted aslong as it retains json format,\n - Lead with ~ to refer to home directory")

  flag.Parse()

  switch {
    case *configPtr != "":
      config_location = strings.Replace(*configPtr, "~", __dir, 1)
    case isEnvSet:
      config_location = strings.Replace(config_env, "~", __dir, 1)
    default:
      config_location = "config.json"
  }

  GENIE.BASE = filepath.Dir(config_location)
  GENIE.Config = config_location

  fmt.Printf("Loading config from: %s\n", GENIE.Config)
}

func main() {
  readflags()

  config := LoadConfig(GENIE.Config)

  selected := PromptUser(config)
  selected.Run()
}
