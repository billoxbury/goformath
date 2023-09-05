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
- profiling using pprof

Run with:

go build .
./gaussian -h
./gaussian -w 2 > output.txt

For profile:

go tool pprof -pdf ./gaussian ./prof/randS.prof > ./prof/randS.pdf
go tool pprof -pdf ./gaussian ./prof/randP.prof > ./prof/randP.pdf
go tool pprof -pdf ./gaussian ./prof/polarS.prof > ./prof/polarS.pdf
go tool pprof -pdf ./gaussian ./prof/polarP.prof > ./prof/polarP.pdf
go tool pprof -pdf ./gaussian ./prof/cartS.prof > ./prof/cartS.pdf
go tool pprof -pdf ./gaussian ./prof/cartP.prof > ./prof/cartP.pdf

etc

*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	f, _ := os.Create("gaussian.prof")
	pprof.StartCPUProfile(f)
	pprof.StopCPUProfile()

	// cmd line arguments
	var n, w int
	flag.IntVar(&n, "n", 1e06, "sample size")
	flag.IntVar(&w, "w", 1, "nr workers")
	flag.Parse()

	// run rnorm generators
	// 1.
	fmt.Println("Serial sample (rand):")

	f1, _ := os.Create("randS.prof")
	pprof.StartCPUProfile(f1)
	start1 := time.Now()
	x1 := rnorm(n, rand.NormFloat64)
	tStd := time.Since(start1)
	pprof.StopCPUProfile()

	fmt.Printf("Time:\t\t%v\n", tStd)
	summary(x1)

	// 2.
	fmt.Println("Parallel sample (rand):")

	f2, _ := os.Create("randP.prof")
	pprof.StartCPUProfile(f2)
	start2 := time.Now()
	x2 := rnormPar(n, w, rand.NormFloat64)
	tStdPar := time.Since(start2)
	pprof.StopCPUProfile()

	fmt.Printf("Time:\t\t%v\n", tStdPar)
	summary(x2)

	// 3.
	fmt.Println("Serial Box-Muller polar:")

	f3, _ := os.Create("polarS.prof")
	pprof.StartCPUProfile(f3)
	start3 := time.Now()
	x3 := rnorm(n, rnormBoxMullerPolar)
	tBMpol := time.Since(start3)
	pprof.StopCPUProfile()

	fmt.Printf("Time:\t\t%v\n", tBMpol)
	summary(x3)

	// 4.
	fmt.Println("Parallel Box-Muller polar:")

	f4, _ := os.Create("polarP.prof")
	pprof.StartCPUProfile(f4)
	start4 := time.Now()
	x4 := rnormPar(n, w, rnormBoxMullerPolar)
	tBMpolPar := time.Since(start4)
	pprof.StopCPUProfile()

	fmt.Printf("Time:\t\t%v\n", tBMpolPar)
	summary(x4)

	// 5.
	fmt.Println("Serial Box-Muller cartesian:")

	f5, _ := os.Create("cartS.prof")
	pprof.StartCPUProfile(f5)
	start5 := time.Now()
	x5 := rnorm(n, rnormBoxMullerCart)
	tBMcart := time.Since(start5)
	pprof.StopCPUProfile()

	fmt.Printf("Time:\t\t%v\n", tBMcart)
	summary(x5)

	// 6.
	fmt.Println("Parallel Box-Muller cartesian:")

	f6, _ := os.Create("cartP.prof")
	pprof.StartCPUProfile(f6)
	start6 := time.Now()
	x6 := rnormPar(n, w, rnormBoxMullerCart)
	tBMcartPar := time.Since(start6)
	pprof.StopCPUProfile()

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
