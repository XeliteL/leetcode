package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow(1/x, -n)
	}

	if n == 0 {
		return 1
	} else if n%2 == 0 {
		half := myPow(x, n/2)
		return half * half
	} else {
		half := myPow(x, (n-1)/2)
		return half * half * x
	}
}
