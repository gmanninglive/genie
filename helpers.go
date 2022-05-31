package main

import "reflect"


type helper interface{}

// #if block helper
func ifHelper() interface{} {
	return "yes"
}


var helpers = make(map[string]reflect.Value)
func Init() { 
	helpers["ifHelper"] = reflect.ValueOf(ifHelper)
}