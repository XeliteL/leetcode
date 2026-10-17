package main

// 23. Merge k Sorted Lists
// https://leetcode.com/problems/merge-k-sorted-lists/

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

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{
			name:     "три списка",
			lists:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:     "пустой массив списков",
			lists:    [][]int{},
			expected: []int{},
		},
		{
			name:     "массив с пустым списком",
			lists:    [][]int{{}},
			expected: []int{},
		},
		{
			name:     "один список",
			lists:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "два списка",
			lists:    [][]int{{1, 3}, {2, 4}},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "списки разной длины",
			lists:    [][]int{{1}, {1, 2, 3}, {2}},
			expected: []int{1, 1, 2, 2, 3},
		},
		{
			name:     "все списки пустые",
			lists:    [][]int{{}, {}, {}},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := make([]*ListNode, len(tt.lists))
			for i, l := range tt.lists {
				nodes[i] = buildList(l)
			}
			result := mergeKLists(nodes)
			got := listToSlice(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("mergeKLists(%v) = %v, ожидалось %v",
					tt.lists, got, tt.expected)
			}
		})
	}
}
