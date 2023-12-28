package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	funcName := "add"
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
	result := reflect.ValueOf(add).Call(args)
	fmt.Printf("%s result is %v\n", funcName, result[0].Int())

	funcName = "sub"
	args = []reflect.Value{reflect.ValueOf(5), reflect.ValueOf(3)}
	result = reflect.ValueOf(sub).Call(args)
	fmt.Printf("%s result is %v\n", funcName, result[0].Int())
}
