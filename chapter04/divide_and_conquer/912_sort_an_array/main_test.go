package main

// 912. Sort an Array
// https://leetcode.com/problems/sort-an-array/

import "testing"

func TestSortArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "случайный порядок",
			nums:     []int{5, 2, 3, 1},
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "с дубликатами",
			nums:     []int{5, 1, 1, 2, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name:     "уже отсортирован",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "обратный порядок",
			nums:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "один элемент",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "отрицательные числа",
			nums:     []int{-5, -2, -3, -1, 0},
			expected: []int{-5, -3, -2, -1, 0},
		},
		{
			name:     "все элементы одинаковые",
			nums:     []int{3, 3, 3, 3},
			expected: []int{3, 3, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortArray(tt.nums)
			if len(result) != len(tt.expected) {
				t.Fatalf("sortArray(%v) длина = %d, ожидалось %d",
					tt.nums, len(result), len(tt.expected))
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("sortArray(%v) = %v, ожидалось %v",
						tt.nums, result, tt.expected)
					break
				}
			}
		})
	}
}
