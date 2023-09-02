/*

- cmd line arguments with flag package
- use of math/rand package


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
> ./coinflip --help
> ./coinflip -n 20 -p 0.2


*/

package main

import (
	"flag"
	"fmt"
)

func main() {

	// cmd line arguments
	var N, M int
	var p float64
	flag.IntVar(&N, "n", 10, "nr coin tosses/binomial draws")
	flag.IntVar(&M, "m", 1000, "nr experiments")
	flag.Float64Var(&p, "p", 0.5, "Bernoulli probability")
	flag.Parse()

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
