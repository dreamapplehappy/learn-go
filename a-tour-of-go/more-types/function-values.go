package main

import (
	"math"
	"fmt"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3.0, 4.0)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot))
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(math.Pow))
}