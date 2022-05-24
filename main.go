package main

import (
	"flag"
	"fmt"
)

type Env struct {
  Config string
  BASE string
}

type Flags struct {
  Config string
}

var GENIE Env

func Check(e error) {
  if e != nil {
      panic(e)
  }
}

func readflags() Flags {
  configPtr := flag.String("config", "", "set location of config.json:\n - Include filename,\n - Custom filnames are accepted aslong as it retains json format,\n - Lead with ~ to refer to home directory")
  
  flag.Parse()

  return Flags{ Config: *configPtr }
}

func main() {
  flags := readflags()

  config := LoadConfig(flags)

  selected := PromptUser(config)
  fmt.Printf("%s\n", selected)
  selected.Run()
}
