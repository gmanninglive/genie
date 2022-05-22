package main

import (
	"flag"
	"fmt"
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

func scheduler(t Task) {
  for i := 0; i < len(t.Schedule); i++ {
    current := t.Schedule[i]

    runCommand(current, t.Base, t.Vars)
  }
}

func runCommand(c Command, __base string, ctx CtxVars) {
  template_path := filepath.Join(__base, c.Template)

  parsed := HandleBars(template_path, ctx)

  c.Output = filepath.Join(c.Directory, c.Filename)

  WriteFile(c, []byte(parsed))
  fmt.Printf("created: %s\n", c.Output)
}

func main() {
  flags := readflags()
  config := LoadConfig(flags.Config)
  fmt.Println(config)

  selected, err := SelectTask(config)
  Check(err)

  withVars := SetVars(config[selected])

  scheduler(withVars)
}
 