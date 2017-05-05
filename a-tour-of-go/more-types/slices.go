package main

import "fmt"

func main() {
	var primes = []int{2, 3, 5, 7, 11, 13}
	s := primes[1:4]
	fmt.Println(s)
}
