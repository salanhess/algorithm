package main

import (
	"fmt"
	. "testbh/SinglyLinkedList"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func deleteDuplicates2(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func main() {
	//SinglyLinkedList.Test()
	s := New()
	s.Format([]int{1, 1, 2, 3, 111, 111, 222})
	//s.Traverse()
	head := s.Front()
	head = deleteDuplicates(head)
	s.Traverse()
	fmt.Println("=======above easy to use ======")
	s = New()
	var n1 *ListNode = &ListNode{1, nil}
	s.Append(n1)
	n2 := &ListNode{2, nil}
	s.Append(n2)
	n22 := &ListNode{2, nil}
	s.Append(n22)
	n3 := &ListNode{3, nil}
	s.Append(n3)
	s.Traverse()
	head = deleteDuplicates2(n1)
	//head = deleteDuplicates(n1)
	s.Traverse()
}
