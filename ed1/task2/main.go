package main

import "fmt"

func gcd(len1, len2 int) int {
	for len2 != 0 {
		len1, len2 = len2, len1%len2
	}
	return len1
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	return str2[:gcd(len(str1), len(str2))]
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))  // ABC
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // AB
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // ""
	fmt.Println(gcdOfStrings("AAAAAB", "AAA"))  // ""
}
