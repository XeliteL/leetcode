package main

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

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

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "обычный список",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "два элемента",
			nums:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "один элемент",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "пустой список",
			nums:     []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.nums)
			result := reverseList(head)
			got := listToSlice(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseList(%v) = %v, ожидалось %v",
					tt.nums, got, tt.expected)
			}
		})
	}
}
