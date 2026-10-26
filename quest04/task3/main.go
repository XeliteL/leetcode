package main

import "fmt"

func largestRectangleArea(heights []int) int {
	res := 0
	stack := make([]int, 0, len(heights))
	n := len(heights)

	for i := 0; i <= n; i++ {
		height := 0
		if i < n {
			height = heights[i]
		}
		for len(stack) != 0 && heights[stack[len(stack)-1]] > height {
			top := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) != 0 {
				width = i - stack[len(stack)-1] - 1
			}
			res = max(res, top*width)
		}

		stack = append(stack, i)
	}

	return res
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // 10
	fmt.Println(largestRectangleArea([]int{2, 4}))             // 4
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))          // 3
}
