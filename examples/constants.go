package main

import (
	"fmt"
	"math"
)

const s string = "good study day day up"

func main() {
	fmt.Println(s)

	// 数值型常量没有确定类型，直到使用时才确定
	const n = 500000

	const d = 3e+20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
