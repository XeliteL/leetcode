package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	res := 0
	target := tickets[k]

	for i, num := range tickets {
		if i <= k {
			res += min(num, target)
		} else {
			res += min(num, target-1)
		}
	}

	return res
}

func main() {
	fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))                                                                                                              // 6
	fmt.Println(timeRequiredToBuy([]int{5, 1, 1, 1}, 0))                                                                                                           // 8
	fmt.Println(timeRequiredToBuy([]int{84, 49, 5, 24, 70, 77, 87, 8}, 3))                                                                                         // 154
	fmt.Println(timeRequiredToBuy([]int{15, 66, 3, 47, 71, 27, 54, 43, 97, 34, 94, 33, 54, 26, 15, 52, 20, 71, 88, 42, 50, 6, 66, 88, 36, 99, 27, 82, 7, 72}, 18)) // 1457
}
