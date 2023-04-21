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

func traverseRecursive(head *ListNode) {
	if head == nil {
		return
	}
	// 前序位置
	traverseRecursive(head.Next)
	// 后序位置
}
