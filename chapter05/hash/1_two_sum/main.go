package main

func twoSum(nums []int, target int) []int {
	idx := map[int]int{}

	for i, num := range nums {
		needNum := target - num
		_, ok := idx[needNum]

		if !ok {
			idx[num] = i
		} else {
			return []int{i, idx[needNum]}
		}
	}

	return nil
}
