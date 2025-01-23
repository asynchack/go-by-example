package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0

	for _, i := range nums {
		total += i
	}

	fmt.Println(total)
	fmt.Println(nums[1:2])
}

func main() {

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
