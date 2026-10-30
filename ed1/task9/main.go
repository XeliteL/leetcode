package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	write := 0

	for read := 0; read < len(chars); {
		ch := chars[read]
		count := 0
		for read < len(chars) && chars[read] == ch {
			read++
			count++
		}
		chars[write] = ch
		write++
		if count > 1 {
			for _, d := range strconv.Itoa(count) {
				chars[write] = byte(d)
				write++
			}
		}
	}

	return write
}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))                               // 6
	fmt.Println(compress([]byte{'a'}))                                                             // 1
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'})) // 4
}
