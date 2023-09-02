package main

import (
	"math/rand"
)

func binomial(N int, p float64) int {

	h := 0
	for i := 0; i < N; i++ {
		if heads(p) {
			h++
		}
	}
	return h
}

func heads(p float64) bool {
	return rand.Float64() < p
}
