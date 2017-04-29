package main

import "fmt"

// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
func tryDefer() int {
	fmt.Println("before defer")
	defer fmt.Println("Hello, World")
	fmt.Println("after defer")
	return 1
}

func main() {
	m := tryDefer()
	fmt.Println(m)
}