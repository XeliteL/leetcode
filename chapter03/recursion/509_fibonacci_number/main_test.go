package main

// 509. Fibonacci Number
// https://leetcode.com/problems/fibonacci-number/

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "F(0) = 0",
			n:        0,
			expected: 0,
		},
		{
			name:     "F(1) = 1",
			n:        1,
			expected: 1,
		},
		{
			name:     "F(2) = 1",
			n:        2,
			expected: 1,
		},
		{
			name:     "F(3) = 2",
			n:        3,
			expected: 2,
		},
		{
			name:     "F(4) = 3",
			n:        4,
			expected: 3,
		},
		{
			name:     "F(10) = 55",
			n:        10,
			expected: 55,
		},
		{
			name:     "F(30) = 832040",
			n:        30,
			expected: 832040,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fib(tt.n)
			if result != tt.expected {
				t.Errorf("fib(%d) = %d, ожидалось %d",
					tt.n, result, tt.expected)
			}
		})
	}
}
