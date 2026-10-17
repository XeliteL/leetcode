package main

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]

		switch {
		case target == guess:
			return mid
		case target > guess:
			low = mid + 1
		case target < guess:
			high = mid - 1
		}
	}

	return -1
}
