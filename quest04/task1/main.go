package main

import "fmt"

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))
	copy(res, prices)
	stack := make([]int, 0, len(prices))

	for i, price := range prices {
		for len(stack) != 0 && prices[stack[len(stack)-1]] >= price {
			res[stack[len(stack)-1]] -= price
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3})) //[4,2,4,2,3]
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5})) //[1,2,3,4,5]
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))   //[9,0,1,6]
}
