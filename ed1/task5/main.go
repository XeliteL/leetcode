package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			i++
		}
	}

	return n == 0
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)) // true
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)) // false
}
