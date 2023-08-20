/*

- command line options and flag package
- recursion
- --help
- usage check and os.Exit()

Run with:

> go build .
> ./gcd 720 1040
> ./gcd -v 720 1040
> ./gcd -h

*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
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

	// compute
	d := gcd(a, b)

	// report
	if verbose {
		fmt.Printf("gcd(%v, %v) = %v\n", a, b, d)
		if d > 1 {
			fmt.Printf("%v = %v x %v\n", a, d, a/d)
			fmt.Printf("%v = %v x %v\n", b, d, b/d)
		}
	} else {
		fmt.Println(d)
	}
	// end
}
