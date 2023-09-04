/*

Tests, for runtime and statistical fit, three ways to generate a univariate normal sample:
using built-in rand.NormFloat64(), using Box-Muller with polar coordinates and using Cartesian Box-Muller
with rejection sampling.

For each of these methods, compares a simple serial implementation with a parallel implementation using goroutines.

A fairly significant speed-up is seen for the rand.float64() when the number of workers
is set to the number of cores of the machine. For Box-Muller the advantage seems to be lost, on the other hand.

- passing function type as argument to another function
- parallelisation with goroutines
- use of a buffered channel
- append slices with ...
- read data from file

Run with:

> go build .
> ./gaussian -h
> ./gaussian -w 2

*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// cmd line arguments
	var n, w int
	flag.IntVar(&n, "n", 1e06, "sample size")
	flag.IntVar(&w, "w", 1, "nr workers")
	flag.Parse()

	// run rnorm generators
	// 1.
	fmt.Println("Serial sample (rand):")
	start1 := time.Now()
	x1 := rnorm(n, rand.NormFloat64)
	tStd := time.Since(start1)
	fmt.Printf("Time:\t\t%v\n", tStd)
	summary(x1)

	// 2.
	fmt.Println("Parallel sample (rand):")
	start2 := time.Now()
	x2 := rnormPar(n, w, rand.NormFloat64)
	tStdPar := time.Since(start2)
	fmt.Printf("Time:\t\t%v\n", tStdPar)
	summary(x2)

	// 3.
	fmt.Println("Serial Box-Muller polar:")
	start3 := time.Now()
	x3 := rnorm(n, rnormBoxMullerPolar)
	tBMpol := time.Since(start3)
	fmt.Printf("Time:\t\t%v\n", tBMpol)
	summary(x3)

	// 4.
	fmt.Println("Parallel Box-Muller polar:")
	start4 := time.Now()
	x4 := rnormPar(n, w, rnormBoxMullerPolar)
	tBMpolPar := time.Since(start4)
	fmt.Printf("Time:\t\t%v\n", tBMpolPar)
	summary(x4)

	// 5.
	fmt.Println("Serial Box-Muller cartesian:")
	start5 := time.Now()
	x5 := rnorm(n, rnormBoxMullerCart)
	tBMcart := time.Since(start5)
	fmt.Printf("Time:\t\t%v\n", tBMcart)
	summary(x5)

	// 6.
	fmt.Println("Parallel Box-Muller cartesian:")
	start6 := time.Now()
	x6 := rnormPar(n, w, rnormBoxMullerCart)
	tBMcartPar := time.Since(start6)
	fmt.Printf("Time:\t\t%v\n", tBMcartPar)
	summary(x6)

}

// summary stats
func summary(x []float64) {
	fmt.Printf("Sample size\t%d\n", len(x))
	fmt.Printf("Mean\t\t%v\n", mean(x))
	fmt.Printf("Variance\t%v\n", variance(x))
	fmt.Printf("Deviance\t%v\n", deviance(x))
}
