package main

func main() {
	sum := 0
	// 在for循环内的初始变量仅在for语句的作用域中可见
	for i := 1; i <= 10; i++ {
		sum += i
	}
	println(sum)
}
