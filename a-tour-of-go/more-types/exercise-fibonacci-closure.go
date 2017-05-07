package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second, fib := 0, 1, 0
	isFirst := true
	return func() int {
		if isFirst {
			isFirst = false
		} else {
			fib = first + second
			first = second
			second = fib
		}
		return fib
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
