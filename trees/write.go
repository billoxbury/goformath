package main

import (
	"fmt"
	"io"
)

// method to present a view of the tree on the screen
func (t *BinaryTree) print(w io.Writer) {

	t.root.printR(w, 0, 'o')
	fmt.Fprintf(w, "\n")
}

func (node *BinaryNode) printR(w io.Writer, ns int, ch rune) {

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	if node == nil {
		fmt.Fprintf(w, "%c: -\n", ch)
		return
	}
	switch node.count {
	case 1:
		fmt.Fprintf(w, "%c: %v\n", ch, node.data)
	default:
		fmt.Fprintf(w, "%c: %v(%d)\n", ch, node.data, node.count)
	}

	node.left.printR(w, ns+2, 'l')
	node.right.printR(w, ns+2, 'r')
}

// method to output Newick string representation
func (t *BinaryTree) newick(w io.Writer) {

	t.root.newickR(w)
	fmt.Fprintf(w, ";\n")
}

func (node *BinaryNode) newickR(w io.Writer) {

	if node == nil {
		fmt.Fprintf(w, "")
		return
	}
	if node.left != nil || node.right != nil {
		fmt.Fprintf(w, "(")
		node.left.newickR(w)
		if node.left != nil && node.right != nil {
			fmt.Fprintf(w, ",")
		}
		node.right.newickR(w)
		fmt.Fprintf(w, ")\n")
	}
	fmt.Fprintf(w, "%v:%d", node.data, node.count)
}

// return words in alphabetical order with counts
func (t *BinaryTree) sorted(w io.Writer) {

	t.root.sortedR(w)
}

func (node *BinaryNode) sortedR(w io.Writer) {

	if node == nil {
		return
	}
	node.left.sortedR(w)
	fmt.Fprintf(w, "%3d %v\n", node.count, node.data)
	node.right.sortedR(w)
}
