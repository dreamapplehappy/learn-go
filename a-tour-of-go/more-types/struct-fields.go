package main

import "fmt"

type Vertex1 struct {
	X int
	Y int
}

func main() {
	v := Vertex1{1, 2}
	// 使用点号.来访问结构体字段
	fmt.Println(v.X)
	v.X = 100
	fmt.Println(v.X)
}
