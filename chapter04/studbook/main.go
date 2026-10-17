package main

import "fmt"

func sum(xd []int) int {
	if len(xd) == 0 {
		return 0
	}

	return xd[0] + sum(xd[1:])
}

func count(xd []int) int {
	if len(xd) == 0 {
		return 0
	}

	return 1 + count(xd[1:])
}

func max(xd []int) int {
	if len(xd) == 0 {
		return -1
	} else if len(xd) == 1 {
		return xd[0]
	} else {
		if xd[0] >= max(xd[1:]) {
			return xd[0]
		} else {
			return max(xd[1:])
		}
	}
}

func binarySearchRecursion(xd []int, need, offset int) int {
	if len(xd) == 0 {
		return -1
	}

	mid := len(xd) / 2

	if need == xd[mid] {
		return mid + offset
	} else if need > xd[mid] {
		return binarySearchRecursion(xd[mid:], need, offset+mid)
	} else {
		return binarySearchRecursion(xd[:mid+1], need, offset)
	}
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(count([]int{1, 2, 3, 4, 5}))
	fmt.Println(max([]int{1, 5, 2, 4, 3}))
	fmt.Println(binarySearchRecursion([]int{1, 2, 3, 4, 5}, 4, 0))
}
