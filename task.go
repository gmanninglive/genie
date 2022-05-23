package main

import (
	"fmt"
	"path/filepath"

	"github.com/aymerick/raymond"
)

type Task struct {
  Title string
  Schedule []Command
  Params []string
  Vars TplVars
  Parser Parser
}

type Command struct {
  Directory string
  Filename string
  Template string
  Output string
}

type TplVars map[string]string

func (t Task) Run() {
  t.Parser.Init()

  for i := 0; i < len(t.Schedule); i++ {
    current := t.Schedule[i]
    current = t.parseSchedule(current)
    t.runCommand(current, t.Vars)
  }
}

func (t Task) runCommand(c Command, tplvars TplVars) {
  template_path := filepath.Join(__base, c.Template)
  
  parsed := t.Parser.Parse(template_path, tplvars)

  c.Output = filepath.Join(c.Directory, c.Filename)

  WriteFile(c, []byte(parsed))
  fmt.Printf("created: %s\n", c.Output)
}

func (t Task) parseSchedule(c Command) Command{
  c.Directory = raymond.MustRender(c.Directory, t.Vars)
  c.Filename = raymond.MustRender(c.Filename, t.Vars)

  return c
}