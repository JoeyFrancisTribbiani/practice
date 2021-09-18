package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	h := new(ListNode)
	h.Next = head
	p := h
	for {
		if p.Next != nil {
			if p.Next.Val == val {
				p.Next = p.Next.Next
			} else {
				p = p.Next
			}
		} else {
			break
		}
	}
	return h.Next
}
