package main

import (
	"fmt"
	"math/rand"
	"time"
)

func naiveSearch(perm []int, dist [][]float64, niters int) ([]int, float64) {

	npoints := len(dist)
	ndescents := 0
	start := time.Now()
	for iter := 0; iter < niters; iter++ {

		i := rand.Intn(npoints)
		j := rand.Intn(npoints)
		delta_d := deltaDist(i, j, perm, dist)
		if delta_d < 0 {
			swap(i, j, perm)
			ndescents += 1
		}
	}
	runtime := time.Since(start)
	distance := travelDist(perm, dist)
	fmt.Printf("Found distance %v in %d steps, time %v\n", distance, ndescents, runtime)

	return perm, distance
}
