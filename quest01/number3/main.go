package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	res, cur := 0, 0

	for _, num := range nums {
		if num == 1 {
			cur++
			res = max(res, cur)
		} else {
			cur = 0
		}
	}

	return res
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
