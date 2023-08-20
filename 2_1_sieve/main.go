/*

- constants (like #define in C)
- command line args
- convert string to int
- initialising a constant-length array

Run with:

> go build .
> ./sieve 50

*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

const MAXN int = 100000

func main() {

	// read argument from the command line
	N, _ := strconv.Atoi(os.Args[1])

	// initialise a sieve table - false by default, true if found composite
	var table [MAXN]bool

	// sieve
	for i := 2; i < N; i++ {
		if !table[i] {
			for j := i; i*j < N; j++ {
				table[i*j] = true
			}
		}
	}

	// read off remaining falses from table
	for i := 2; i < N; i++ {
		if !table[i] {
			fmt.Printf("%4v\n", i)
		}
	}
}
