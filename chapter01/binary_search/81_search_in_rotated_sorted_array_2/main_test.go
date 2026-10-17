package main

// 81. Search in Rotated Sorted Array II
// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

import "testing"

func TestSearchRotatedII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "target найден — дубликаты есть",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   0,
			expected: true,
		},
		{
			name:     "target не найден — дубликаты есть",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   3,
			expected: false,
		},
		{
			name:     "массив из одного элемента — найден",
			nums:     []int{1},
			target:   1,
			expected: true,
		},
		{
			name:     "массив из одного элемента — не найден",
			nums:     []int{1},
			target:   0,
			expected: false,
		},
		{
			name:     "все элементы одинаковые — найден",
			nums:     []int{1, 1, 1, 1, 1},
			target:   1,
			expected: true,
		},
		{
			name:     "все элементы одинаковые — не найден",
			nums:     []int{1, 1, 1, 1, 1},
			target:   2,
			expected: false,
		},
		{
			name:     "дубликаты на границах",
			nums:     []int{1, 1, 1, 0, 1},
			target:   0,
			expected: true,
		},
		{
			name:     "массив не повёрнут — найден",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: true,
		},
		{
			name:     "массив не повёрнут — не найден",
			nums:     []int{1, 2, 3, 4, 5},
			target:   6,
			expected: false,
		},
		{
			name:     "дубликаты скрывают точку поворота",
			nums:     []int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			target:   5,
			expected: true,
		},
		{
			name:     "target на границе поворота",
			nums:     []int{3, 1, 1},
			target:   3,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := search(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("search(%v, %d) = %v, ожидалось %v",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
