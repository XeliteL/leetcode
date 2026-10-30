package main

import (
	"fmt"
)

func maxOperations(nums []int, k int) int {
	res := 0
	cnt := make(map[int]int, len(nums)-1)

	for _, num := range nums {
		if cnt[k-num] > 0 {
			res++
			cnt[k-num]--
		} else {
			cnt[num]++
		}
	}

	return res
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))                                           // 2
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))                                        // 1
	fmt.Println(maxOperations([]int{2, 1, 3, 4}, 5))                                           // 2
	fmt.Println(maxOperations([]int{3, 5, 1, 5}, 2))                                           // 0
	fmt.Println(maxOperations([]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2)) // 2
}
