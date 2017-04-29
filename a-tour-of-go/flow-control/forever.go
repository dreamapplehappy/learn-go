package main

import "fmt"

func main() {
	sum := 1
	// 如果省略循环条件，并且循环体内没有终止循环的条件，该循环就不会结束
	for {
		// 后面是自己添加的
		fmt.Println(sum)
		sum += sum
		if sum > 10 {
			break;
		}
	}
}
