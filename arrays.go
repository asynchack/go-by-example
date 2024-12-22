package main

import (
	"fmt"
)

func main()  {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	fmt.Println(a[4])

	b := [...]int{1, 2, 3}
	fmt.Println(b)
	fmt.Println(len(b))

	//idx 5是6，0和5之间的1-4 是类型默认值
	var c = [...]int{1, 5: 66, 77}
	fmt.Println(len(c)) // 7
	fmt.Println(c) 

	var td [2][3]int
	for i:=0; i < 2; i++ {
		for j:=0; j < 3; j++ {
			td[i][j] = i+j
		}
	}
	fmt.Println(td)

	td2 := [2][3]int{
		{1, 1, 1},
		{2, 2, 2},
	}
	fmt.Println(td2)
}