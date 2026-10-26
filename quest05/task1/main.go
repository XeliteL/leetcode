package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	count := make([]int, 2)

	for _, s := range students {
		count[s]++
	}

	for i, s := range sandwiches {
		if count[s] == 0 {
			return len(sandwiches) - i
		}
		count[s]--
	}

	return 0
}

func main() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))             //0
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1})) //3
}
