package main

// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "пример 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "один элемент",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "k равен числу уникальных элементов",
			nums:     []int{1, 2},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "несколько с одинаковой частотой",
			nums:     []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "xD",
			nums:     []int{5, 3, 1, 1, 1, 3, 5, 73, 1},
			k:        3,
			expected: []int{1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)
			sort.Ints(result)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("topKFrequent(%v, %d) = %v, ожидалось %v",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
