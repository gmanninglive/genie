package helpers

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {

	Init()

	funcVal := helpers["ifHelper"]

	// argType := funcVal.Type().In(0)

	// for i := range args {
	// 	if canBeNil(argType) {
	// 		args[i] = reflect.Zero(argType)
	// 	}
	// }
	req := reflect.TypeOf(funcVal).NumIn()

	//args := reflect.ValueOf("string")
	//res := reflect.ValueOf(funcVal).Call([]reflect.Value{args})
	//t.Log(res)
	t.Log(req)
}

func canBeNil(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}
