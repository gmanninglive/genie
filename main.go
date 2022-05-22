package main

import (
	"flag"
	"fmt"
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
  fmt.Println(config)

  selected := PromptUser(config)

  selected.InitParser()
  selected.Run()
}
 