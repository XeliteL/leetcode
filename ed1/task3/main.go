package main

import (
	"fmt"
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxCandies := slices.Max(candies)

	for i := range candies {
		res[i] = candies[i]+extraCandies >= maxCandies
	}

	return res
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)) // [true true true false true]
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)) // [true false false false false]
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))    // [true false true]
}
