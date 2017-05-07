package main

import "fmt"

type Vertex5 struct {
	Lat, Long float64
}

var m1 = map[string]Vertex5{
	"dream": Vertex5{1.234, 3.456},
	"happy": Vertex5{6.777, 8.999},
}

func main() {
	fmt.Println(m1)
}

