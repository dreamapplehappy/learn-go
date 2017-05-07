package main

import "fmt"

type Vertex4 struct {
	Lat, Long float64
}

// 这里是声明
var m map[string]Vertex4

func main() {
	m = make(map[string]Vertex4)
	m["dreamapple "] = Vertex4{12.345, 34.667}
	fmt.Println(m["dreamapple "])
}
