package main

import (
	"time"
	"fmt"
)

func main() {
	hour := time.Now().Hour()
	// 没有条件的switch语句；这种形式能将一长串 if-then-else 写得更加清晰
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
