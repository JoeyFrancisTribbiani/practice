package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h1, h2 := head, head.Next
	i, j := h1, h2
	for {
		if i.Next == nil || j.Next == nil {
			break
		}
		if i.Next.Next != nil {
			i.Next = i.Next.Next
			i = i.Next
		}
		j.Next = j.Next.Next
		j = j.Next
	}
	i.Next = h2
	return h1
}
