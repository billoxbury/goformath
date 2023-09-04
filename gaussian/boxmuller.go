package main

import (
	"math"
	"math/rand"
)

// Box-Muller generator in polar coordinates
func rnormBoxMullerPolar() float64 {

	s := -2 * math.Log(rand.Float64())
	theta := 2 * math.Pi * rand.Float64()
	return math.Sqrt(s) * math.Cos(theta)
}

// rejection-sampled Box-Muller in cartesian coordinates
func rnormBoxMullerCart() float64 {

	for {
		x := -1 + 2*rand.Float64()
		y := -1 + 2*rand.Float64()
		s := x*x + y*y
		if s <= 1 {
			return x * math.Sqrt(-2*math.Log(s)/s)
		}
	}
}
