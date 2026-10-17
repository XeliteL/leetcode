package main

// 215. Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "пример из условия 1",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "пример из условия 2",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "k=1, ищем максимум",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        1,
			expected: 6,
		},
		{
			name:     "k равен длине массива, ищем минимум",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        6,
			expected: 1,
		},
		{
			name:     "один элемент",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "все элементы одинаковые",
			nums:     []int{2, 2, 2, 2},
			k:        3,
			expected: 2,
		},
		{
			name:     "отрицательные числа",
			nums:     []int{-1, -5, -3, -2, -4},
			k:        2,
			expected: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("findKthLargest(%v, %d) = %d, ожидалось %d",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
