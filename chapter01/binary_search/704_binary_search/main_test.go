package main

// 704. Binary Search
// https://leetcode.com/problems/binary-search/

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "target найден в середине",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "target найден в начале",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   -1,
			expected: 0,
		},
		{
			name:     "target найден в конце",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   12,
			expected: 5,
		},
		{
			name:     "target не найден",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			name:     "массив из одного элемента - найден",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "массив из одного элемента - не найден",
			nums:     []int{5},
			target:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("search(%v, %d) = %d, ожидалось %d",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
