package main

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}

func partition(nums []int, low, high int) (int, int) {
	lo := low
	mid := low
	hi := high
	xdNum := nums[lo]

	for mid <= hi {
		if nums[mid] < xdNum {
			nums[lo], nums[mid] = nums[mid], nums[lo]
			lo++
			mid++
		} else if nums[mid] == xdNum {
			mid++
		} else {
			nums[mid], nums[hi] = nums[hi], nums[mid]
			hi--
		}
	}

	return lo, hi
}

func quicksort(nums []int, low, high int) {
	if low < high {
		lo, hi := partition(nums, low, high)
		quicksort(nums, low, lo-1)
		quicksort(nums, hi+1, high)
	}
}
