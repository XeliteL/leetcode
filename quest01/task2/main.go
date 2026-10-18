package main

import "fmt"

func shuffle(nums []int, n int) []int {
	res := make([]int, 0, n*2)

	for i := range n {
		res = append(res, nums[i], nums[n+i])
	}

	return res
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	fmt.Println(shuffle([]int{1, 1, 2, 2}, 2))
}
