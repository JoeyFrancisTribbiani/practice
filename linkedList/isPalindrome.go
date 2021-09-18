package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head.Next
	var t *ListNode = nil
	var r *ListNode = nil
	for {
		t = slow
		slow = slow.Next
		t.Next = r
		r = t
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		if fast == nil {
			slow = slow.Next
			break
		}
	}
	for slow != nil {
		if slow.Val != r.Val {
			return false
		}
		slow = slow.Next
		r = r.Next
	}
	return true
}
