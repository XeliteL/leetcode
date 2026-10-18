package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		if num < 0 {
			num = -num
		}

		if nums[num-1] > 0 {
			nums[num-1] *= -1
		}
	}

	res := nums[:0]

	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
