/*

See
https://amehta.github.io/posts/2019/03/wc-implementation-in-go-lang/

- error handling
- os.Open() and file.Close()
- bufio scanner

Run with:

> go build .
> ./wordcount main.go


*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// MAIN
func main() {

	// read file name
	var fileName string

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		fmt.Printf("Usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}

	// open file
	file, err := os.Open(fileName)
	check(err)

	// create scanner (bufio)
	scanner := bufio.NewScanner(file)

	// initialise counts
	lines, words, characters := 0, 0, 0

	// run scan
	for scanner.Scan() {
		lines++

		line := scanner.Text()
		characters += len(line)

		splitLines := strings.Split(line, " ")
		words += len(splitLines)
	}

	// report
	fmt.Printf("%8d%8d%8d\t%s\n", lines, words, characters, fileName)

	// close file
	file.Close()
}
