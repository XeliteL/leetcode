package main

var pick int

func guess(num int) int {
	if num > pick {
		return -1
	}
	if num < pick {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	low := 1
	high := n

	for low <= high {
		num := (low + high) / 2

		switch guess(num) {
		case -1:
			high = num - 1
		case 0:
			return num
		case 1:
			low = num + 1
		}
	}

	return -1
}
