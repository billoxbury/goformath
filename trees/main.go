/*

Aim:

- make interface and implementation for binary trees
- demonstrate independent client apps that use this package
- include drawing among tree methods

Adapted from
https://www.golangprograms.com/golang-program-to-implement-binary-tree.html

Covers:

- switch
- range
- runes vs strings
- structs and methods
- io.Writer

Run with:

go build .
./trees -h

text=`cat br.txt | tail -n 26 | head -n 7 | tr '[:upper:]' '[:lower:]' | tr -d ",|(|)|?|'" `

echo $text | ./trees -p
echo $text | ./trees -s

echo $text | ./trees > newick.txt
Rscript ./drawTree.R newick.txt

// --> open Rplots.pdf

*/

package main

import (
	"flag"
	"os"
)

// specify the item type for nodes
type itemType string

func main() {

	// read file name
	var file *os.File

	// cmd line arguments
	var filename string
	var prnt bool
	var sort bool
	flag.StringVar(&filename, "f", "", "[optional] input file name")
	flag.BoolVar(&prnt, "p", false, "print schematically [default: Newick string]")
	flag.BoolVar(&sort, "s", false, "return sorted word list [default: Newick string]")
	flag.Parse()

	// open file or stdin
	if filename != "" {
		file, _ = os.Open(filename)
	} else {
		file = os.Stdin
	}

	tree := &BinaryTree{}
	tree.read(file)
	if prnt {
		tree.print(os.Stdout)
	} else if sort {
		tree.sorted(os.Stdout)
	} else {
		tree.newick(os.Stdout)
	}
}
