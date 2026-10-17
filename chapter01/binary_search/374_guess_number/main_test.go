package main

// 374. Guess Number Higher or Lower
// https://leetcode.com/problems/guess-number-higher-or-lower/

import "testing"

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		pick     int
		expected int
	}{
		{
			name:     "число в середине",
			n:        10,
			pick:     6,
			expected: 6,
		},
		{
			name:     "единственное число",
			n:        1,
			pick:     1,
			expected: 1,
		},
		{
			name:     "число в начале диапазона",
			n:        10,
			pick:     1,
			expected: 1,
		},
		{
			name:     "число в конце диапазона",
			n:        10,
			pick:     10,
			expected: 10,
		},
		{
			name:     "два числа — загадано первое",
			n:        2,
			pick:     1,
			expected: 1,
		},
		{
			name:     "два числа — загадано второе",
			n:        2,
			pick:     2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pick = tt.pick
			result := guessNumber(tt.n)
			if result != tt.expected {
				t.Errorf("guessNumber(%d) = %d, ожидалось %d",
					tt.n, result, tt.expected)
			}
		})
	}
}
