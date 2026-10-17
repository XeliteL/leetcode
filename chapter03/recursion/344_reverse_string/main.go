package main

func reverseString(s []byte) {
	rec(0, len(s)-1, s)
}

func rec(left, right int, s []byte) {
	if left > right {
		return
	}
	s[left], s[right] = s[right], s[left]
	left++
	right--
	rec(left, right, s)
}
