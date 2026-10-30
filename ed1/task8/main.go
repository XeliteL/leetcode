package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else if num > second && second > first {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))    // true
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))    // false
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6})) // true
}
