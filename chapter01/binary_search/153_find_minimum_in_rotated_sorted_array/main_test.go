package main

// 153. Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "минимум в середине",
			nums:     []int{3, 4, 5, 1, 2},
			expected: 1,
		},
		{
			name:     "минимум в центре",
			nums:     []int{3, 1, 2},
			expected: 1,
		},
		{
			name:     "минимум после большого блока",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			expected: 0,
		},
		{
			name:     "массив не повёрнут",
			nums:     []int{11, 13, 15, 17},
			expected: 11,
		},
		{
			name:     "массив из одного элемента",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "массив из двух элементов — повёрнут",
			nums:     []int{2, 1},
			expected: 1,
		},
		{
			name:     "массив из двух элементов — не повёрнут",
			nums:     []int{1, 2},
			expected: 1,
		},
		{
			name:     "повёрнут на одну позицию",
			nums:     []int{2, 3, 4, 5, 1},
			expected: 1,
		},
		{
			name:     "повёрнут на n позиций (полный оборот)",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMin(tt.nums)
			if result != tt.expected {
				t.Errorf("findMin(%v) = %d, ожидалось %d",
					tt.nums, result, tt.expected)
			}
		})
	}
}
