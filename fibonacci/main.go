package main

import (
	"fmt"
)

func main() {

	// get index from user
	fmt.Printf("Up to: ")
	var n, idx int
	fmt.Scanln(&n)
	idx = n + 1

	// declare and make a slice
	var arr []int
	arr = make([]int, idx)

	// THE RIGHT WAY
	fib(arr)
	for i := 0; i < idx; i++ {
		fmt.Printf("%3d %12d\n", i, arr[i])
	}

	// THE WRONG WAY
	fmt.Println("Recursive implementation:")
	fmt.Printf("%3d %12d\n", n, fibNaive(n))
}
