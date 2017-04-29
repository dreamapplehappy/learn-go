package main

import "fmt"

func main() {
	// sum 初始值要为一个正数
	sum := 1
	// for的初始语句和后置语句是可选的
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
