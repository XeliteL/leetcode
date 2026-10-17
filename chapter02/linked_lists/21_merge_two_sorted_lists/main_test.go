package main

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/

import (
	"reflect"
	"testing"
)

// сборка списка из слайса для удобства тестов
func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

// список в слайс для удобства сравнения в тестах
func listToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "оба непустые, чередуются",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "оба пустые",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "первый пустой",
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
		{
			name:     "второй пустой",
			list1:    []int{1, 2, 3},
			list2:    []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "list1 полностью меньше list2",
			list1:    []int{1, 2, 3},
			list2:    []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "list2 полностью меньше list1",
			list1:    []int{4, 5, 6},
			list2:    []int{1, 2, 3},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "с дубликатами между списками",
			list1:    []int{1, 1, 1},
			list2:    []int{1, 1, 1},
			expected: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name:     "значения на границах диапазона",
			list1:    []int{-100, 0, 100},
			list2:    []int{-100, 0, 100},
			expected: []int{-100, -100, 0, 0, 100, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := buildList(tt.list1)
			l2 := buildList(tt.list2)
			result := mergeTwoLists(l1, l2)
			got := listToSlice(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("mergeTwoLists(%v, %v) = %v, ожидалось %v",
					tt.list1, tt.list2, got, tt.expected)
			}
		})
	}
}
