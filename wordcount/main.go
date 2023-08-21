/*

See
https://amehta.github.io/posts/2019/03/wc-implementation-in-go-lang/

- open file or stdio
- bufio scanner
- strings package

NOTE char count still needs correcting!

Run with:

> go build .
> ./wordcount go.mod
> ./wordcount main.go
> wc main.go

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// read file name
	var fileName string
	var file *os.File

	// open file or stdin
	if len(os.Args) > 1 {
		fileName = os.Args[1]
		file, _ = os.Open(fileName)
	} else {
		fileName = ""
		file = os.Stdin
	}

	// create scanner (bufio)
	scanner := bufio.NewScanner(file)

	// initialise counts
	lines, words, characters := 0, 0, 0

	// run scan
	for scanner.Scan() {
		lines++

		line := scanner.Text()
		characters += len(line)

		// choose a splitting function ...
		//splitLines := strings.Split(line, " ") // doesn't ignore empty lines
		splitLines := strings.Fields(line)
		words += len(splitLines)
	}

	// report
	fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, fileName)

	// close file
	file.Close()
}
