/*

- bit ops
- timing
- %v string format for complex 'value' (used here for time)

Run with:

> go build .
> ./bingcd 720 1040
> ./bingcd -v 720 1040


*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	// simple usage check
	n := len(os.Args)
	if n < 2 {
		fmt.Printf("Usage: %s [-v] n1 n2\n", os.Args[0])
		os.Exit(1)
	}

	// cmd line flags
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.Parse()

	// cmd line args (must come after any flags)
	a, _ := strconv.Atoi(os.Args[n-2])
	b, _ := strconv.Atoi(os.Args[n-1])

	// compute Euclidean
	start := time.Now()
	d1 := gcdEuc(a, b, verbose)
	t1 := time.Since(start)
	if verbose {
		fmt.Println("-------------------------------------------")
	}

	// compute binary
	start = time.Now()
	d2 := gcdBin(a, b, verbose)
	t2 := time.Since(start)
	if verbose {
		fmt.Println("-------------------------------------------")
	}

	// report
	fmt.Printf("Euclidean: \t%d\t%v\n", d1, t1)
	fmt.Printf("Binary: \t%d\t%v\n", d2, t2)

	// end
}
