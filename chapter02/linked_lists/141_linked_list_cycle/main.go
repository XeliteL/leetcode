package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	currFast := head
	currSlow := head

	if head == nil {
		return false
	}

	for currFast.Next != nil && currFast.Next.Next != nil {
		currFast = currFast.Next.Next
		currSlow = currSlow.Next

		if currSlow == currFast {
			return true
		}
	}

	return false
}
