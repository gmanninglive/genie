package main

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {

	Init()

	funcVal := helpers["ifHelper"]
	
	// argType := funcVal.Type().In(0)
	args := make([]reflect.Value, 0)

	// for i := range args {
	// 	if canBeNil(argType) {
	// 		args[i] = reflect.Zero(argType)
	// 	}
	// }
	
	t.Log(funcVal.Call(args))
}

func canBeNil(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}