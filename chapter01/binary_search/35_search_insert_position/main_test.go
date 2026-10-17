package main

// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "target найден",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "target вставляется в середину",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "target вставляется в конец",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "target вставляется в начало",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "target найден в начале",
			nums:     []int{1, 3, 5, 6},
			target:   1,
			expected: 0,
		},
		{
			name:     "массив из одного элемента — найден",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "массив из одного элемента — вставка перед",
			nums:     []int{1},
			target:   0,
			expected: 0,
		},
		{
			name:     "массив из одного элемента — вставка после",
			nums:     []int{1},
			target:   2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchInsert(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("searchInsert(%v, %d) = %d, ожидалось %d",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
