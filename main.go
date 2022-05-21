package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"
	"github.com/manifoldco/promptui"
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

type MakeCMD struct {
  Directory string
  Filename string
  Template string
  Output string
}

type Command struct {
  Title string
  Make []MakeCMD
  CtxVars map[string]string
  Base string
}

type Config []Command

func loadConfig(config_location string) Config {
  var config Config
  
  f, err := os.ReadFile(config_location)
  check(err)
  
  json.Unmarshal([]byte(f), &config)
  
  for i := 0; i < len(config); i++ {
    config[i].Base = filepath.Dir(config_location)
  }
  
  return config
}

func runCommand(command Command) {
  for i := 0; i < len(command.Make); i++ {
    template_path := filepath.Join(command.Base, command.Make[i].Template)

    parse, err := raymond.ParseFile(template_path)
    check(err)
    withCtx, err := parse.Exec(command.CtxVars)
    check(err)

    command.Make[i].Output = filepath.Join(command.Make[i].Directory, command.Make[i].Filename)

    writeFile(command.Make[i], []byte(withCtx))
    fmt.Printf("created: %s\n", command.Make[i].Output)
  }
}

func writeFile(make MakeCMD, template []byte) {
  if _, err := os.Stat(make.Directory); os.IsNotExist(err) {
    err := os.MkdirAll(make.Directory, 0700)
    check(err)
  }

  err := os.WriteFile(make.Output, []byte(template), 0644)
  check(err)
}

func main() {
  flags := readflags()
  config := loadConfig(flags.Config)
  options := make([]string, len(config))

  for i := 0; i < len(config); i++ {
    options[i] = config[i].Title
  }

  prompt := promptui.Select{
		Label: "Select a Command",
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
  
  for i := 0; i < len(config); i++ {
    if config[i].Title == result {
      runCommand(config[i])

      return
    }
  }
}
 