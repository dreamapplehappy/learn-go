package main

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func main() {
	v := Vertex2{1, 3}
	p := &v
	// 结构体指针可以使用隐式间接引用 不需要使用(*p)表示
	p.X = 100
	fmt.Println(v)
	v1 := v
	v1.X = 3
	fmt.Println(v1.X)
}
