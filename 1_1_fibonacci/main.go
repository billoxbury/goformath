/*

- get data from stdin
- make an array (length read from stdin)
- function call (written in a separate file)
- for loop
- string formatting

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

	// declare array
	var arr []int
	arr = make([]int, idx)

	// compute
	fib(arr)

	// report
	for i := 0; i < idx; i++ {
		fmt.Printf("%3v %12v\n", i, arr[i])
	}
}
