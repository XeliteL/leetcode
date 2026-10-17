package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMid(head)
	second := mid.Next
	mid.Next = nil

	left := sortList(head)
	right := sortList(second)

	return mergeTwoLists(left, right)
}

func findMid(head *ListNode) *ListNode {
	fastCurr := head
	slowCurr := head

	for fastCurr.Next != nil && fastCurr.Next.Next != nil {
		fastCurr = fastCurr.Next.Next
		slowCurr = slowCurr.Next
	}

	return slowCurr
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		xd := mergeTwoLists(list1.Next, list2)
		list1.Next = xd
		return list1
	} else {
		xd := mergeTwoLists(list2.Next, list1)
		list2.Next = xd
		return list2
	}
}
