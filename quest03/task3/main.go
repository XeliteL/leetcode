package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	stack := make([]int, 0, len(logs)/2)

	prev := 0
	for _, log := range logs {
		fVh := strings.IndexByte(log, ':')
		sVh := strings.IndexByte(log[fVh+1:], ':') + fVh + 1
		ts, _ := strconv.Atoi(log[sVh+1:])
		if log[fVh+1] == 's' {
			idx, _ := strconv.Atoi(log[:fVh])
			if len(stack) != 0 {
				res[stack[len(stack)-1]] += ts - prev
			}
			stack = append(stack, idx)
			prev = ts
		} else {
			res[stack[len(stack)-1]] += ts - prev + 1
			stack = stack[:len(stack)-1]
			prev = ts + 1
		}
	}

	return res
}

func main() {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))                         //[3,4]
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"})) //[8]
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"})) //[7,1]
}
