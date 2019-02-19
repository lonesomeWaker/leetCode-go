package swapPairs

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	h := head
	if h == nil {
		return h
	}
	swap := h.Next
	if swap == nil {
		return h
	}
	nextSwap := swap.Next
	swap.Next = h
	h.Next = swapPairs(nextSwap)
	return swap
}
