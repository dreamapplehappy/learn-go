package main

import "fmt"

type Vertex6 struct {
	Lat, Long float64
}

var m2 = map[string]Vertex6{
	"a": {1.0, 2.0},
}

func main() {
	fmt.Println(m2)
}
