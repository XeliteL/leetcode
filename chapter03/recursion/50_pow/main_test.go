package main

// 50. Pow(x, n)
// https://leetcode.com/problems/powx-n/

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{
			name:     "положительная степень",
			x:        2.00000,
			n:        10,
			expected: 1024.00000,
		},
		{
			name:     "дробное основание",
			x:        2.10000,
			n:        3,
			expected: 9.26100,
		},
		{
			name:     "отрицательная степень",
			x:        2.00000,
			n:        -2,
			expected: 0.25000,
		},
		{
			name:     "степень 0",
			x:        5.00000,
			n:        0,
			expected: 1.00000,
		},
		{
			name:     "степень 1",
			x:        5.00000,
			n:        1,
			expected: 5.00000,
		},
		{
			name:     "основание 1",
			x:        1.00000,
			n:        1000,
			expected: 1.00000,
		},
		{
			name:     "основание -1, чётная степень",
			x:        -1.00000,
			n:        4,
			expected: 1.00000,
		},
		{
			name:     "основание -1, нечётная степень",
			x:        -1.00000,
			n:        3,
			expected: -1.00000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := myPow(tt.x, tt.n)
			if math.Abs(result-tt.expected) > 1e-5 {
				t.Errorf("myPow(%v, %d) = %v, ожидалось %v",
					tt.x, tt.n, result, tt.expected)
			}
		})
	}
}
