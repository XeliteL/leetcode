package main

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	index := 0

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]
		index = mid

		switch {
		case guess == target:
			return mid
		case guess < target:
			low = mid + 1
			index++
		case guess > target:
			high = mid - 1
		}
	}

	return index
}
