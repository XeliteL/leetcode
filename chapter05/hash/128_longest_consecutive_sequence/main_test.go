package main

// 128. Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "пример 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "пример 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "дубликаты",
			nums:     []int{1, 0, 1, 2},
			expected: 3,
		},
		{
			name:     "пустой массив",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "один элемент",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "все элементы одинаковые",
			nums:     []int{3, 3, 3},
			expected: 1,
		},
		{
			name:     "отрицательные числа",
			nums:     []int{-3, -2, -1, 0, 1},
			expected: 5,
		},
		{
			name:     "две отдельные последовательности",
			nums:     []int{1, 2, 3, 10, 11, 12, 13},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestConsecutive(tt.nums)
			if result != tt.expected {
				t.Errorf("longestConsecutive(%v) = %d, ожидалось %d",
					tt.nums, result, tt.expected)
			}
		})
	}
}
