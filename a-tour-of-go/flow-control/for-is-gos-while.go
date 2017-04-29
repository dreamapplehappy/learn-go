package main

import "fmt"

func main() {
	sum := 1
	for sum < 1e4 {
		sum += sum
	}
	fmt.Println(sum)
}
