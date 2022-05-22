package main

import (
	"fmt"
	"path/filepath"
)

func (t Task) InitParser() {
  var p Parser
  p.Init()
  t.parser = p
}

func (t Task) Run() {
  for i := 0; i < len(t.Schedule); i++ {
    current := t.Schedule[i]

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
