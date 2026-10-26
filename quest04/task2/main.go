package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for i, temperature := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperature {
			top := stack[len(stack)-1]
			res[top] = i - top
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) //[1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                 //[1,1,1,0]
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))                     //[1,1,0]
}
