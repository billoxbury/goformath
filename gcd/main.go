/*

- command line options and flag package
- strconv package
- recursion
- --help
- usage check and os.Exit()

Run with:

> go build .
> ./gcd -h
> ./gcd 720 1040
> ./gcd -x 720 1040

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
	var explain bool
	flag.BoolVar(&explain, "x", false, "explain")
	flag.Parse()

	// cmd line args (must come after any flags)
	a, _ := strconv.Atoi(os.Args[n-2])
	b, _ := strconv.Atoi(os.Args[n-1])

	// compute
	d := gcd(a, b)

	// report
	if explain {
		fmt.Printf("gcd(%d, %d) = %d\n", a, b, d)
		if d > 1 {
			fmt.Printf("%d = %d x %d\n", a, d, a/d)
			fmt.Printf("%d = %d x %d\n", b, d, b/d)
		}
	} else {
		fmt.Println(d)
	}
	// end
}
