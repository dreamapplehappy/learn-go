package main

import "fmt"

// type用来声明类型
// struct用来声明结构体
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
