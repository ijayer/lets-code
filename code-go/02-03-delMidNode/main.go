package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	deleteNode(node.Next.Next.Next.Next)

	for node.Next != nil {
		fmt.Printf("node: %+v\n", *node)

		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
	fmt.Printf("node: %+v\n", *node)
}

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	if node.Next != nil {
		next := node.Next
		node.Val = next.Val
		node.Next = next.Next
	}
}
