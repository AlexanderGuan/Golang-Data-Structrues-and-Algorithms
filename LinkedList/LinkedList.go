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

// 递归反转链表前N个节点
var successor *ListNode

// 定义：反转以 head 为起点的 n 个节点，返回新的头结点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n + 1 个节点
		successor = head.Next
		return head
	}
	// 以 head.Next 为起点，需要反转前 n - 1 个节点
	last := reverseN(head.Next, n-1)

	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = successor
	return last
}
