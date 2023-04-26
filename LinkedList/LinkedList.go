package linkedlist

import "fmt"

type ListNode struct {
	Val  int64
	Next *ListNode
}

// 迭代遍历链表
func traverse(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

// 递归遍历链表
func traverseRecursive(head *ListNode) {
	if head == nil {
		return
	}
	// 前序位置
	traverseRecursive(head.Next)
	// 后序位置
}

// 递归反转整个单链表
// 定义：输入一个单链表头结点，将该链表反转，返回新的头结点
func reverseRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
