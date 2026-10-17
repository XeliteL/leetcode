package main

func containsDuplicate(nums []int) bool {
	contain := map[int]bool{}

	for _, v := range nums {
		if contain[v] {
			return true
		}
		contain[v] = true
	}

	return false
}
