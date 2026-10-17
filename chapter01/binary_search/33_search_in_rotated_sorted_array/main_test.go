package main

// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "target найден во второй части",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "target не найден",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "массив из одного элемента — не найден",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "массив из одного элемента — найден",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "target найден в первой части",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		{
			name:     "target первый элемент",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   4,
			expected: 0,
		},
		{
			name:     "target последний элемент",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   2,
			expected: 6,
		},
		{
			name:     "массив не повёрнут",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "массив не повёрнут — target не найден",
			nums:     []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "два элемента — повёрнут",
			nums:     []int{3, 1},
			target:   1,
			expected: 1,
		},
		{
			name:     "три элемента — target последний",
			nums:     []int{5, 1, 3},
			target:   3,
			expected: 2,
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
