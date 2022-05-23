package main

import (
	"flag"
	"path/filepath"
)

type Flags struct {
  Config string
}

func readflags() Flags {
  configPtr := flag.String("config", "config.json", "location of config.json")
  flag.Parse()

  res := Flags{ Config : *configPtr }

  return res
}

func main() {
  flags := readflags()
  config := LoadConfig(flags.Config)

  selected := PromptUser(config)
  selected.Base = filepath.Dir(flags.Config)
  selected.Run()
}
