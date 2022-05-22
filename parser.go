package main

import (
	"github.com/aymerick/raymond"
)

func HandleBars(template_path string, vars map[string]string) string {
  parse, err := raymond.ParseFile(template_path)

  Check(err)
  withCtx, err := parse.Exec(vars)
  Check(err)

  return withCtx
}
