package main

import (
	"fmt"
	"path/filepath"

	"github.com/aymerick/raymond"
)

func (t Task) Run() {
  t.parser.Init()

  for i := 0; i < len(t.Schedule); i++ {
    current := t.Schedule[i]
    current = t.parseSchedule(current)
    t.runCommand(current, t.Base, t.Vars)
  }
}

func (t Task) runCommand(c Command, __base string, ctx CtxVars) {
  template_path := filepath.Join(__base, c.Template)
  
  parsed := t.parser.Parse(template_path, ctx)

  c.Output = filepath.Join(c.Directory, c.Filename)

  WriteFile(c, []byte(parsed))
  fmt.Printf("created: %s\n", c.Output)
}

func (t Task) parseSchedule(c Command) Command{
  c.Directory = raymond.MustRender(c.Directory, t.Vars)
  c.Filename = raymond.MustRender(c.Filename, t.Vars)

  return c
}