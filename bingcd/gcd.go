package main

import "fmt"

func gcdEuc(a int, b int, verbose bool) int {

	if verbose {
		fmt.Printf("%8d\t%8d\n", a, b)
	}
	if b == 0 {
		return a
	} else {
		return gcdEuc(b, a%b, verbose)
	}
}

func gcdBin(a int, b int, verbose bool) int {

	// extract common power of 2
	k := 0
	for a&1 == 0 && b&1 == 0 {
		a >>= 1
		b >>= 1
		k++
	}

	// initialise t to be even if possible
	var t int
	if a&1 == 1 {
		t = -b
	} else {
		t = a
	}

	// extract powers of 2 from t;
	// replace max(a,b) with t;
	// take t to be the new difference;
	// repeat until t==0
	if verbose {
		fmt.Printf("%8d\t%8d\t%8d\n", t, a, b)
	}
	for t != 0 {
		for t&1 == 0 {
			t >>= 1
		}
		if t > 0 {
			a = t
		} else {
			b = -t
		}
		t = a - b
		if verbose {
			fmt.Printf("%8d\t%8d\t%8d\n", t, a, b)
		}
	}
	// replace the common power of 2
	return a << k
}
