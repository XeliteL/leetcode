package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	res := make([]int, 2)

	for _, num := range nums {
		if num < 0 {
			num = -num
		}

		if nums[num-1] < 0 {
			res[0] = num
		} else {
			nums[num-1] *= -1
		}
	}

	for i := range nums {
		if nums[i] > 0 {
			res[1] = i + 1
			return res
		}
	}

	return res
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
}
