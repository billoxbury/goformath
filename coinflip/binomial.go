package main

import (
	"math/rand"
)

func binomial(N int, p float64) int {

	h := 0
	for i := 0; i < N; i++ {
		if rand.Float64() < p {
			h++
		}
	}
	return h
}
