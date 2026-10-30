package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	f := 0

	for _, symb := range t {
		if s[f] == byte(symb) {
			f++
			if len(s) == f {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
}
