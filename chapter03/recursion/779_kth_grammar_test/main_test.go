package main

// 779. K-th Symbol in Grammar
// https://leetcode.com/problems/k-th-symbol-in-grammar/

import "testing"

func TestKthGrammar(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected int
	}{
		{
			name:     "строка 1, позиция 1",
			n:        1,
			k:        1,
			expected: 0,
		},
		{
			name:     "строка 2, позиция 1",
			n:        2,
			k:        1,
			expected: 0,
		},
		{
			name:     "строка 2, позиция 2",
			n:        2,
			k:        2,
			expected: 1,
		},
		{
			name:     "строка 3, позиция 1",
			n:        3,
			k:        1,
			expected: 0,
		},
		{
			name:     "строка 3, позиция 2",
			n:        3,
			k:        2,
			expected: 1,
		},
		{
			name:     "строка 3, позиция 3",
			n:        3,
			k:        3,
			expected: 1,
		},
		{
			name:     "строка 3, позиция 4",
			n:        3,
			k:        4,
			expected: 0,
		},
		{
			name:     "строка 4, позиция 5",
			n:        4,
			k:        5,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kthGrammar(tt.n, tt.k)
			if result != tt.expected {
				t.Errorf("kthGrammar(%d, %d) = %d, ожидалось %d",
					tt.n, tt.k, result, tt.expected)
			}
		})
	}
}
