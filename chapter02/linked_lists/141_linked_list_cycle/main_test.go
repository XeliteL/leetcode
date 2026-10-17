package main

// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/

import "testing"

// buildListWithCycle строит список из nums и замыкает последний узел
// на узел с индексом pos (с нуля). pos = -1 означает "без цикла".
func buildListWithCycle(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	curr := head
	nodes := []*ListNode{head}

	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
		nodes = append(nodes, curr)
	}

	if pos != -1 {
		curr.Next = nodes[pos]
	}

	return head
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		pos      int
		expected bool
	}{
		{
			name:     "цикл есть, замыкается на второй узел",
			nums:     []int{3, 2, 0, -4},
			pos:      1,
			expected: true,
		},
		{
			name:     "цикл есть, замыкается на себя (два узла)",
			nums:     []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			name:     "цикла нет",
			nums:     []int{1},
			pos:      -1,
			expected: false,
		},
		{
			name:     "цикла нет, длинный список",
			nums:     []int{1, 2, 3, 4, 5},
			pos:      -1,
			expected: false,
		},
		{
			name:     "пустой список",
			nums:     []int{},
			pos:      -1,
			expected: false,
		},
		{
			name:     "цикл замыкается на первый узел (весь список — цикл)",
			nums:     []int{1, 2, 3},
			pos:      0,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildListWithCycle(tt.nums, tt.pos)
			result := hasCycle(head)
			if result != tt.expected {
				t.Errorf("hasCycle(%v, pos=%d) = %v, ожидалось %v",
					tt.nums, tt.pos, result, tt.expected)
			}
		})
	}
}
