package main

import (
	"fmt"
	"io"
)

// method to present a view of the tree on the screen
func (node *BinaryNode) print(w io.Writer, ns int, ch rune) {

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

	(node.left).print(w, ns+2, 'l')
	(node.right).print(w, ns+2, 'r')
}
