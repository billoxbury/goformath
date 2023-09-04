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
	start1 := time.Now()
	x1 := rnorm(n, rand.NormFloat64)
	tStd := time.Since(start1)
	fmt.Printf("Serial sample (rand) size %d\n", len(x1))

	start2 := time.Now()
	x2 := rnormPar(n, w, rand.NormFloat64)
	tStdPar := time.Since(start2)
	fmt.Printf("%d combined samples (rand) total %d\n", w, len(x2))

	start3 := time.Now()
	x3 := rnorm(n, rnormBoxMullerPolar)
	tBMpol := time.Since(start3)
	fmt.Printf("Serial sample (BM polar) size %d\n", len(x3))

	start4 := time.Now()
	x4 := rnormPar(n, w, rnormBoxMullerPolar)
	tBMpolPar := time.Since(start4)
	fmt.Printf("%d combined samples (BM polar) total %d\n", w, len(x4))

	start5 := time.Now()
	x5 := rnorm(n, rnormBoxMullerCart)
	tBMcart := time.Since(start5)
	fmt.Printf("Serial sample (BM cartesian) size %d\n", len(x5))

	start6 := time.Now()
	x6 := rnormPar(n, w, rnormBoxMullerCart)
	tBMcartPar := time.Since(start6)
	fmt.Printf("%d combined samples (BM cartesian) total %d\n", w, len(x6))

	// report timings
	fmt.Printf("Rand:\t\t\t%v\n", tStd)
	fmt.Printf("Rand (par):\t\t%v\n", tStdPar)
	fmt.Printf("BM polar:\t\t%v\n", tBMpol)
	fmt.Printf("BM polar (par):\t\t%v\n", tBMpolPar)
	fmt.Printf("BM cartesian:\t\t%v\n", tBMcart)
	fmt.Printf("BM cartesian (par):\t%v\n", tBMcartPar)

	fmt.Println(normalQuantiles())

}
