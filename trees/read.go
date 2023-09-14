package main

import (
	"bufio"
	"os"
	"strings"
)

// method to read tree from file or stdin
// NOTE this assumes itemType is string
func (t *BinaryTree) read(file *os.File) {

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		words := strings.Fields(line)
		for _, w := range words {
			t.insert(itemType(w)) // <------- itemType must be string
		}
	}
}
