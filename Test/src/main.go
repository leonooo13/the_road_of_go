package main

import (
	"fmt"
	"runtime"
	"time"
)

func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func swap(a *int, b *int) {
	// 修改地址上的值
	*a, *b = *b, *a
	// tmp := *a
	// *a = *b
	// *b = tmp

}
func swap1(a, b int) (int, int) {
	// 修改地址上的值
	// *a, *b = *b, *a
	return b, a

}

func arry() {

	var slice []int

	slice = make([]int, 0, 10)
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	for i, v := range slice {
		fmt.Println(i, v)
	}
	len_slice := len(slice)
	slice = append(slice, 3, 4, 5)
	for i := len_slice; i < len(slice); i++ {
		fmt.Println(i, slice[i])
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	// go arry()
	// go loop()

	time.Sleep(time.Second)

	a := 1
	b := 2
	fmt.Println(a, b)
	fmt.Println(&a, &b)
	a, b = swap1(a, b)
	// swap1和swap2本质上都是一样的修改地址上的值
	fmt.Println(&a, &b)
	fmt.Println(a, b)
	//修改值
	c := 10
	fmt.Println(&c)
	c = 20
	fmt.Println(&c)

}
