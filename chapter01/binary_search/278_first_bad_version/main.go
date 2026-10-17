package main

var firstBad int

func isBadVersion(version int) bool {
	return version >= firstBad
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	first := 1

	for low <= high {
		mid := (low + high) / 2

		switch isBadVersion(mid) {
		case true:
			first = mid
			high = mid - 1
		case false:
			low = mid + 1
		}
	}

	return first
}
