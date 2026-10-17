package main

func kthGrammar(n, k int) int {
	if k == 1 {
		return 0
	} else {
		xd := kthGrammar(n-1, (k+1)/2)
		if k%2 == 1 {
			return xd
		} else {
			if xd == 1 {
				return 0
			}
			return 1
		}
	}
}
