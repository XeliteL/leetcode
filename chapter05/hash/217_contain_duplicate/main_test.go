package main

// 217. Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "есть дубликат",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "нет дубликатов",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "несколько дубликатов",
			nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{
			name:     "один элемент",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "два одинаковых элемента",
			nums:     []int{5, 5},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("containsDuplicate(%v) = %v, ожидалось %v",
					tt.nums, result, tt.expected)
			}
		})
	}
}
