package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabet := map[rune]int{}

	for _, v := range s {
		alphabet[v]++
	}

	for _, v := range t {
		alphabet[v]--
		if alphabet[v] < 0 {
			return false
		}
	}

	for _, v := range alphabet {
		if v != 0 {
			return false
		}
	}

	return true
}
