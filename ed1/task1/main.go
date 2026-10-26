package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	res := make([]byte, len(word1)+len(word2))
	maxLen := max(len(word1), len(word2))

	for i := range maxLen {
		if i > len(word1) {
			res = append(res, word1[i])
		}
		if i > len(word2) {
			res = append(res, word2[i])
		}
	}

	return string(res)
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr")) // apbqcr
	fmt.Println(mergeAlternately("ab", "pqrs")) // apbqrs
	fmt.Println(mergeAlternately("abcd", "pq")) // apbqcd
}
