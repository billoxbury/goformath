/*

Aim:

- make interface and implementation for binary trees
- demonstrate independent client apps that use this package
- include drawing among tree methods

Adapted from
https://www.golangprograms.com/golang-program-to-implement-binary-tree.html

Covers:

- switch
- structs and methods
- io.Writer

*/

package main

import (
	"os"
)

// specify the item type for nodes
// type itemType int64
type itemType string

// read in payload
// var payload = []itemType{6, 2, 8, 1, 0, 3, 7, 8, 5, 7, 6, 2, 1, 7}
var payload = []itemType{"i", "see", "a", "little", "siloetto", "of", "a", "man", "i", "see", "a", "man"}

func main() {

	tree := &BinaryTree{}

	for _, v := range payload {
		tree.insert(v)
	}

	(tree.root).print(os.Stdout, 0, 'M')
}
