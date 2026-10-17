package main

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	set := map[int]bool{}
	for _, v := range nums {
		set[v] = true
	}

	max := 1
	for num := range set {
		if set[num-1] {
			continue
		}

		length := 1
		for set[num+length] {
			length++
		}

		if length > max {
			max = length
		}
	}

	return max
}
