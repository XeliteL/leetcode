package main

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	res := make([]int, n*2)

	for i, num := range nums {
		res[i], res[n+i] = num, num
	}

	return res
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
	fmt.Println(getConcatenation([]int{1, 3, 2, 1}))
}
