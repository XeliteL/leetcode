package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	mapStrings := map[string][]string{}

	for _, v := range strs {
		r := []rune(v)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		key := string(r)

		mapStrings[key] = append(mapStrings[key], v)
	}

	for _, v := range mapStrings {
		res = append(res, v)
	}

	return res
}
