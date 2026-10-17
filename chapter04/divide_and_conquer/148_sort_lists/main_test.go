package main

// 148. Sort List
// https://leetcode.com/problems/sort-list/

import (
	"reflect"
	"testing"
)

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestSortList(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "обычный список",
			nums:     []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "с отрицательными числами",
			nums:     []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
		},
		{
			name:     "пустой список",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "один элемент",
			nums:     []int{1},
			expected: []int{1},
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
			name:     "с дубликатами",
			nums:     []int{3, 1, 2, 1, 3},
			expected: []int{1, 1, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.nums)
			result := sortList(head)
			got := listToSlice(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("sortList(%v) = %v, ожидалось %v",
					tt.nums, got, tt.expected)
			}
		})
	}
}
