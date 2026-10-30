package main

import (
	"fmt"
)

func res(left, right int, height []int) int {
	return (right - left) * min(height[left], height[right])
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	best := 0

	for left < right {
		best = max(best, res(left, right, height))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return best
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
	fmt.Println(maxArea([]int{1, 2, 1}))                   // 2
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))     // 17
}
