package SinglyLinkedList

// refer to https://blog.csdn.net/wangshubo1989/article/details/71515933

import (
	"fmt"
)

//声明节点
type ListNode struct {
	Val  int //Value interface{}
	next *ListNode
}

type LinkedList struct {
	front  *ListNode
	length int
}

func (n *ListNode) Next() *ListNode {
	return n.next
}

//初始化
func (s *LinkedList) Init() *LinkedList {
	s.length = 0
	return s
}

//New一个链表
func New() *LinkedList {
	return new(LinkedList).Init()
}

//返回头结点
func (s *LinkedList) Front() *ListNode {
	return s.front
}

//返回尾结点
func (s *LinkedList) Back() *ListNode {
	currentNode := s.front
	for currentNode != nil && currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode
}

//添加尾结点
func (s *LinkedList) Append(n *ListNode) {
	if s.front == nil {
		s.front = n
	} else {
		currentNode := s.front
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = n
	}
	s.length++
}

//遍历节点
func (s *LinkedList) Traverse() {
	if s.front == nil {
		fmt.Println("Empty LinkedList")
	} else {
		mylist := []int{}
		currentNode := s.front
		for currentNode.next != nil {
			mylist = append(mylist, currentNode.Val)
			currentNode = currentNode.next
		}
		mylist = append(mylist, s.Back().Val)
		fmt.Println(mylist)
	}
}

//添加头结点
func (s *LinkedList) Prepend(n *ListNode) {
	if s.front == nil {
		s.front = n
	} else {
		n.next = s.front
		s.front = n
	}
	s.length++
}

//在指定结点前添加结点
func (s *LinkedList) InsertBefore(insert *ListNode, before *ListNode) {
	if s.front == before {
		insert.next = before
		s.front = insert
		s.length++
	} else {
		currentNode := s.front
		for currentNode != nil &&
			currentNode.next != nil &&
			currentNode.next != before {
			currentNode = currentNode.next
		}
		if currentNode.next == before {
			insert.next = before
			currentNode.next = insert
			s.length++
		}
	}
}

//在指定结点后添加结点
func (s *LinkedList) InsertAfter(insert *ListNode, after *ListNode) {
	currentNode := s.front
	for currentNode != nil &&
		currentNode != after &&
		currentNode.next != nil {
		currentNode = currentNode.next
	}
	if currentNode == after {
		insert.next = after.next
		after.next = insert
		s.length++
	}
}

//删除指定结点
func (s *LinkedList) Remove(n *ListNode) {
	if s.front == n {
		s.front = n.next
		s.length--
	} else {
		currentNode := s.front

		// search for node n
		for currentNode != nil &&
			currentNode.next != nil &&
			currentNode.next != n {
			currentNode = currentNode.next
		}
		// see if current's next node is n
		// if it's not n, then node n wasn't found in list s
		if currentNode.next == n {
			currentNode.next = currentNode.next.next
			s.length--
		}
	}
}

//init node list for leetcode
func (s *LinkedList) Format(list []int) {
	var n *ListNode
	for _, i := range list {
		n = &ListNode{i, nil}
		s.Append(n)
	}
	s.Traverse()
}

func Test() {
	fmt.Println("hello")
	s := New()
	var n1 *ListNode = &ListNode{1, nil}
	s.Append(n1)
	n2 := &ListNode{2, nil}
	s.Append(n2)
	n3 := &ListNode{3, nil}
	s.Append(n3)
	n0 := &ListNode{0, nil}
	s.Prepend(n0)
	n4 := &ListNode{4, nil}
	s.InsertBefore(n4, n0)
	n5 := &ListNode{5, nil}
	s.InsertAfter(n5, n3)
	s.Traverse()
	fmt.Printf("LinkedList is %v \nLinkedList.front %v ,LinkedList.back%v \nlength %d \n", s, s.Front, s.Back, s.length)
	fmt.Println("==============")
	s = New()
	s.Format([]int{1, 2, 3, 1, 111, 222})
	//fmt.Printf("LinkedList2 length %d,next val is %d\nnext->next val %d\n ", s.length, s.front.Next().Val, s.front.Next().Next().Val)
	fmt.Printf("LinkedList2 length %d,next val is %d\nnext->next val %d\nnext->next-next val %d \nnext->next->next->next val %d \n", s.length, s.front.Next().Val, s.front.Next().Next().Val, s.front.Next().Next().Next().Val, s.front.Next().Next().Next().Next().Val)
}
