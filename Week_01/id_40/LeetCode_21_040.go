package main

import (
	"fmt"
	. "testbh/SinglyLinkedList"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//递归操作，只要保证每个listnode.Next为最小的即可
//终止条件（因为默认排好序的，只要有空的，就将其后所有节点挂置最后即可）
//判断条件：递归，如果l1大就l1.Next, 反之亦然，下个节点为res.Next
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val <= l2.Val {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	} else {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	}
	return res
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
	head = deleteDuplicates(n1)
	s.Traverse()
	fmt.Println("========")
	s = New()
	s.Format([]int{1, 2, 3})
	s2 := New()
	s2.Format([]int{1, 2, 3, 4, 5, 6})
	mergehead := mergeTwoLists(s.Front(), s2.Front())
	merge := New()
	merge.Append(mergehead)
	merge.Traverse()
	fmt.Println("=======above leetcode 21 ======")

}
