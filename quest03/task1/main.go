package main

import "fmt"

func buildArray(target []int, n int) []string {
	res := make([]string, 0, 2*target[len(target)-1])
	targCurr := 0

	for i := 1; i <= n && targCurr < len(target); i++ {
		if i == target[targCurr] {
			res = append(res, "Push")
			targCurr++
		} else {
			res = append(res, "Push", "Pop")
		}
	}

	return res
}

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
	fmt.Println(buildArray([]int{2, 3, 4}, 4))
}
