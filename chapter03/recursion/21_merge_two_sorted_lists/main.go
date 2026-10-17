package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else {
		if list1.Val <= list2.Val {
			xd := mergeTwoLists(list1.Next, list2)
			list1.Next = xd
			return list1
		} else {
			xd := mergeTwoLists(list1, list2.Next)
			list2.Next = xd
			return list2
		}
	}
}
