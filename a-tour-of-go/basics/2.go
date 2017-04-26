package main

// 我们使用分组形式的导入语句
import (
	"fmt"
	"math"
)

func main() {
	// 这里使用了math的方法Sqrt；并且使用了%g占位符
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
