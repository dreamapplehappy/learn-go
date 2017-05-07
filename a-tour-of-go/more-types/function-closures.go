package main

import "fmt"

func adder() func(num int) int {
	sum := 0
	return func(num int) int {
		return sum + num;
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-i))
	}
}
