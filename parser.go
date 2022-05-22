package main

import (
	"strconv"
	"strings"

	"github.com/aymerick/raymond"
)

type Parser struct {}

func (p Parser) Parse(template_path string, vars CtxVars) string {
  parse, err := raymond.ParseFile(template_path)

  Check(err)
  withCtx, err := parse.Exec(vars)
  Check(err)

  return withCtx
}

func (p Parser) Init() {
  raymond.RegisterHelper("add", func(val1, val2 int) string {
      return strconv.Itoa(val1 + val2)
  })

  raymond.RegisterHelper("toUpper", func(val string) string {
      return strings.ToUpper(val)
  })

  raymond.RegisterHelper("toLower", func(val string) string {
      return strings.ToLower(val)
  })

  raymond.RegisterHelper("toTitle", func(val string) string {
      return strings.Title(strings.ToLower(val))
  })
}
