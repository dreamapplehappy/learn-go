package main

import "fmt"

type Vertex3 struct {
	X, Y int
}

func main() {
	v1 := Vertex3{1, 2}
	v2 := Vertex3{X: 1}
	v3 := Vertex3{}
	p := &Vertex3{2, 3}
	fmt.Println(v1, v2, v3, p)
}
