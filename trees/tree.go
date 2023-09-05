package main

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  itemType // <-- itemType specification postponed to client
	count int
}

type BinaryTree struct {
	root *BinaryNode
}
