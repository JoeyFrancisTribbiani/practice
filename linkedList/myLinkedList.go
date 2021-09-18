package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Next *ListNode
	Val  int
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for {
		slow = slow.Next
		if slow == nil {
			return false
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		if fast == slow {
			return true
		}
	}
}
func hasCycleSingle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p := head
	for p.Next != nil {
		if p.Val == 22223 {
			return true
		}
		p.Val = 22223
		p = p.Next
	}
	return false
}

/** Initialize your data structure here. */
func Constructor() ListNode {
	head := new(ListNode)
	head.Next = nil
	head.Val = 0
	return *head
}

/** Get the Value of the index-th node in the linked list. If the index is inValid, return -1. */
func (this *ListNode) Get(index int) int {
	p := this.Next
	for i := 0; i < index; i++ {
		if p.Next == nil && i < index-1 {
			return -1
		}
		p = p.Next
	}
	if p == nil {
		return -1
	}
	return p.Val
}
func (this *ListNode) PrintList() {
	str := "head"
	p := this.Next
	for p != nil {
		if p != nil {
			str += (" -> " + strconv.Itoa(p.Val))
			p = p.Next
		} else {
			str += " -> nil"
			break
		}
	}
	fmt.Println(str)
}

/** Add a node of Value Val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *ListNode) AddAtHead(Val int) {
	p := new(ListNode)
	p.Val = Val
	p.Next = this.Next
	this.Next = p
}

/** Append a node of Value Val to the last element of the linked list. */
func (this *ListNode) AddAtTail(Val int) {
	p := this
	for p.Next != nil {
		p = p.Next
	}
	tail := new(ListNode)
	tail.Val = Val
	tail.Next = nil
	p.Next = tail
}

/** Add a node of Value Val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *ListNode) AddAtIndex(index int, Val int) {
	p := this
	for i := 0; i < index; i++ {
		p = p.Next
		if p == nil && i < index-1 {
			return
		}
	}
	v := new(ListNode)
	v.Val = Val
	v.Next = p.Next
	p.Next = v
}

/** Delete the index-th node in the linked list, if the index is Valid. */
func (this *ListNode) DeleteAtIndex(index int) {
	p := this
	for i := 0; i < index; i++ {
		p = p.Next
		if p == nil && i < index-1 {
			return
		}
	}
	if p == nil || p.Next == nil {
		return
	}
	p.Next = p.Next.Next
}

/**
 * Your ListNode object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(Val);
 * obj.AddAtTail(Val);
 * obj.AddAtIndex(index,Val);
 * obj.DeleteAtIndex(index);
 */
