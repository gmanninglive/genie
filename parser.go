package main

import (
	"strconv"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/iancoleman/strcase"
)

type Parser struct {}

func (p Parser) Parse(template_path string, vars TplVars) string {
  parse, err := raymond.ParseFile(template_path)

  Check(err)
  withCtx, err := parse.Exec(vars)
  Check(err)

  return withCtx
}

func (p Parser) Init() {
  helpers := map[string]interface{} {
    "add": func(val1, val2 int) string { return strconv.Itoa(val1 + val2) },
    "toUpper": func(val string) string { return strings.ToUpper(val) },
    "toLower": func(val string) string { return strings.ToLower(val) },
    "toTitle": func(val string) string { return strings.Title(strings.ToLower(val)) },
    "toCamel": func(val string) string { return strcase.ToCamel(val) },
    "conCat": func(val1 string, val2 string) string { res := val1 + val2; return res },
  }

  raymond.RegisterHelpers(helpers)
}
