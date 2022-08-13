package main

import (
	"testing"
)

func PrintValue(data int) {
	println(data)
}

func TestTree(t *testing.T) {
	binaryTree := &BinaryTree{}
	binaryTree.InsertItem(10)
	binaryTree.InsertItem(2)
	binaryTree.InsertItem(36)
	binaryTree.InsertItem(4)
	binaryTree.InsertItem(52)
	binaryTree.SearchItem(10)
	binaryTree.InorderTraversal(binaryTree.root, PrintValue)
	binaryTree.PreorderTraversal(binaryTree.root, PrintValue)
	binaryTree.PostorderTraversal(binaryTree.root, PrintValue)
}
