package main

// 278. First Bad Version
// https://leetcode.com/problems/first-bad-version/

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		firstBad int
		expected int
	}{
		{
			name:     "плохая версия в середине",
			n:        5,
			firstBad: 4,
			expected: 4,
		},
		{
			name:     "первая версия плохая",
			n:        5,
			firstBad: 1,
			expected: 1,
		},
		{
			name:     "последняя версия плохая",
			n:        5,
			firstBad: 5,
			expected: 5,
		},
		{
			name:     "одна версия и она плохая",
			n:        1,
			firstBad: 1,
			expected: 1,
		},
		{
			name:     "плохая версия вторая из двух",
			n:        2,
			firstBad: 2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstBad = tt.firstBad
			result := firstBadVersion(tt.n)
			if result != tt.expected {
				t.Errorf("firstBadVersion(%d) = %d, ожидалось %d",
					tt.n, result, tt.expected)
			}
		})
	}
}
