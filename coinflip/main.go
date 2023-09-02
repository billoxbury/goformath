/*

math.rand package. Note:

- rand.Seed() is deprecated (see documentation)

Not used here:

- rand.Intn()			discrete nonnegative uniform random
- rand.Shuffle()		random permutation of a list
- rand.Perm()			random permutation of 1,...,n
- rand.ExpFloat64()  	exponential distribution
- rand.NormFloat64()	normal distribution

Run with:

> go build .
> ./coinflip 20 1000 0.2


*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// usage check
	n := len(os.Args)
	if n < 4 {
		fmt.Printf("Usage: %s nr_coins nr_experiments coin_prob\n", os.Args[0])
		os.Exit(1)
	}

	// cmd line args
	N, _ := strconv.Atoi(os.Args[1])
	M, _ := strconv.Atoi(os.Args[2])
	p, _ := strconv.ParseFloat(os.Args[3], 64)

	// initialise score table
	var table []int
	for i := 0; i <= N; i++ {
		table = append(table, 0)
	}

	// run experiments
	for j := 0; j < M; j++ {
		table[binomial(N, p)]++
	}

	// detect start of nonzero counts
	start := 0
	for j := 0; j <= N; j++ {
		if table[j] > 0 {
			start = j
			break
		}
	}
	// detect end of nonzero counts
	end := N
	for j := N; j >= 0; j-- {
		if table[j] > 0 {
			end = j
			break
		}
	}

	// report results
	fmt.Println("")
	for j := start - 1; j <= end+1; j++ {
		if j < 0 || j > N {
			continue
		}
		fmt.Printf("%3d ", j)
		for i := 0; i < table[j]; i += 10 {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
	fmt.Println("")
}
