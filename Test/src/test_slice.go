package main

import (
	"fmt"
	"reflect"
)

func slice() {
	var arr []int
	fmt.Println(nil == arr)
	arr = append(arr, 1)
	arr1 := make([]int, 2)

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("arr1: %v\n", arr1)
}
func slice1() {
	arr := []int{1, 2, 3, 4, 5}
	arr1 := [3]int{1, 2, 3}
	fmt.Printf(reflect.TypeOf(arr).String())
	fmt.Printf(reflect.TypeOf(arr1).String())
	arr = append(arr, 1)
	// arr1 = append(arr1, 1)

}

func main() {
	slice1()

}
