package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	counts := make([]int, 101)

	for _, num := range nums {
		counts[num]++
	}

	sum := 0
	for i := range counts {
		count := counts[i]
		counts[i] = sum
		sum += count
	}

	for i := range nums {
		res[i] = counts[nums[i]]
	}

	return res
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))

}
