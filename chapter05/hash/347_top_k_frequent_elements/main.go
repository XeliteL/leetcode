package main

import (
	"cmp"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	frequency := map[int]int{}

	for _, v := range nums {
		frequency[v]++
	}

	pairs := [][]int{}
	for num := range frequency {
		pairs = append(pairs, []int{num, frequency[num]})
	}

	slices.SortFunc(pairs, func(a, b []int) int {
		return cmp.Compare(b[1], a[1])
	})

	for i := 0; i < k; i++ {
		res = append(res, pairs[i][0])
	}

	return res
}
