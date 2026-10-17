package main

// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/

import (
	"reflect"
	"sort"
	"testing"
)

func normalize(groups [][]string) [][]string {
	for i := range groups {
		sort.Strings(groups[i])
	}
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "пример 1",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}},
		},
		{
			name:     "одна пустая строка",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "одна строка",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "все строки одинаковые",
			strs:     []string{"ab", "ab", "ab"},
			expected: [][]string{{"ab", "ab", "ab"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.strs)
			got := normalize(result)
			exp := normalize(tt.expected)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("groupAnagrams(%v) = %v, ожидалось %v",
					tt.strs, got, exp)
			}
		})
	}
}
