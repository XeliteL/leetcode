package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	b := []byte(s)
	isVowel := func(c byte) bool { return strings.IndexByte("aeiouAEIOU", c) >= 0 }

	for l, r := 0, len(b)-1; l < r; {
		switch {
		case !isVowel(b[l]):
			l++
		case !isVowel(b[r]):
			r--
		default:
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}

	return string(b)
}

func main() {
	fmt.Println(reverseVowels("IceCreAm")) // "AceCreIm"
	fmt.Println(reverseVowels("leetcode")) // "leotcede"
}
