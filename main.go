package main

import (
	"encoding/json"
	"flag"
	"os"
  "fmt"
	"path/filepath"

	"github.com/aymerick/raymond"
)

func check(e error) {
  if e != nil {
      panic(e)
  }
}

type Flags struct {
  Config string
}

func readflags() Flags {
  configPtr := flag.String("config", "config.json", "location of config.json")
  flag.Parse()

  res := Flags{ Config : *configPtr }

  return res
}

type Command struct {
  Title string
  Directory string
  Filename string
  Template string
  CtxVars map[string]string
  Output string
}

type Config []Command

func loadConfig(config_location string) Config {
  var config Config	
  
  f, err := os.ReadFile(config_location)
  check(err)

  json.Unmarshal([]byte(f), &config)

  for i := 0; i < len(config); i++ {
    config[i].Output = filepath.Join(config[i].Directory, config[i].Filename)
  }
  
  return config
}

func runCommand(command Command) string {
  parse, err := raymond.ParseFile(command.Template)
  withCtx, err := parse.Exec(command.CtxVars)
  check(err)

  return withCtx
}

func writeFile(command Command, template []byte) {
  if _, err := os.Stat(command.Directory); os.IsNotExist(err) {
    err := os.MkdirAll(command.Directory, 0700)
    check(err)
  }

  err := os.WriteFile(command.Output, []byte(template), 0644)
  check(err)
}

func main() {
  flags := readflags()
  config := loadConfig(flags.Config)
  
  for i := 0; i < len(config); i++ {
    output := runCommand(config[i])
    writeFile(config[i], []byte(output))
    fmt.Printf("%s completed, created: %s\n", config[i].Title, config[i].Output)
  }
}
 