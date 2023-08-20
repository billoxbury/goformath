package main

import (
	"math"
)

const N_INIT_PRIMES int = 4

var firstPrimes = [...]int{2, 3, 5, 7} // [...] says size of array is fixed
var wheel = [...]int{4, 2, 4, 2, 4, 6, 2, 6}

func get_next_factor(n int) int {

	// we'll need square root of n to bound factor tests
	root_n := math.Sqrt(float64(n))

	// check the initial primes
	for j := 0; j < N_INIT_PRIMES; j++ {
		if n%firstPrimes[j] == 0 {
			return firstPrimes[j]
		}
	}

	// check wheel additions (i.e. avoiding obvious composites) up to sqrt(n)
	for j, d := 0, firstPrimes[N_INIT_PRIMES-1]+wheel[0]; float64(d) < root_n; j, d = (j+1)%len(wheel), d+wheel[j] {

		if n%d == 0 {
			return d
		}
	}

	// else n is prime
	return n
}

func find_all_factors(n int) []int {

	var factors = []int{} // slice type

	for n > 1 {
		f := get_next_factor(n)
		factors = append(factors, f) // append to a slice
		n /= f
	}

	return factors
}
