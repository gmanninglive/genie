package helpers

import (
	"strconv"
	"strings"

	"github.com/ettle/strcase"
)

type helper interface{}

func Init() map[string]interface{} {
	helpers := map[string]interface{}{
		"add":     func(val1, val2 int) string { return strconv.Itoa(val1 + val2) },
		"toUpper": func(val string) string { return strings.ToUpper(val) },
		"toLower": func(val string) string { return strings.ToLower(val) },
		"toTitle": func(val string) string { return strings.Title(strings.ToLower(val)) },
		"toCamel": func(val string) string { return strcase.ToCamel(val) },
		"conCat":  func(val1 string, val2 string) string { res := val1 + val2; return res },
	}
	return helpers
}
