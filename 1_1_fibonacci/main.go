/*

- get data from stdin
- make a slice (length read from stdin)
- function call (written in a separate file)
- for loop
- string formatting a la C printf()

Run with:

> go mod init example.com/fibinacci
> go build .
> ./fibonacci

*/

package main

import (
	"fmt"
)

func main() {

	// get index from user
	fmt.Printf("Up to: ")
	var idx int
	fmt.Scanln(&idx)
	idx++

	// declare and make a slice
	var arr []int
	arr = make([]int, idx)

	// compute
	fib(arr)

	// report
	for i := 0; i < idx; i++ {
		fmt.Printf("%3d %12d\n", i, arr[i])
	}
}
