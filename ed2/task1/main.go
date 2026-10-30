package main

import "fmt"

func moveZeroes(nums []int) {
	l := 0
	for r := range nums {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}

	fmt.Println(nums)
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12}) // [1,3,12,0,0]
	moveZeroes([]int{0})              // [0]
}
