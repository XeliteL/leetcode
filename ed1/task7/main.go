package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	lCurr := 1
	for i := range nums {
		res[i] = lCurr
		lCurr *= nums[i]
	}

	rCurr := 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= rCurr
		rCurr *= nums[i]
	}

	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      // [24,12,8,6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) // [0,0,9,0,0]
}
