package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	var res []string
	curr := -1

	for i, symb := range s {
		if symb == ' ' {
			if curr != -1 {
				res = append(res, s[curr:i])
				curr = -1
			}
			continue
		}

		if curr == -1 {
			curr = i
		}
	}
	if curr != -1 {
		res = append(res, s[curr:])
	}

	slices.Reverse(res)

	return strings.Join(res, " ")
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))  // "blue is sky the"
	fmt.Println(reverseWords("  hello world  "))  // "world hello"
	fmt.Println(reverseWords("a good   example")) // "example good a"
}
