package main

import (
	"fmt"
	"math"
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

func metropolisSearch(perm []int, dist [][]float64, niters int) ([]int, float64) {

	npoints := len(dist)

	// annealing parameters
	temperature := 10.0
	cooling := 0.999
	period := 10000

	// track progress
	acceptance := 0
	this_d := travelDist(perm, dist)
	best_d := travelDist(perm, dist)
	best_p := perm

	start := time.Now()
	for iter := 0; iter < niters; iter++ {

		i := rand.Intn(npoints)
		j := rand.Intn(npoints)
		delta_d := deltaDist(i, j, perm, dist)
		if delta_d < 0 || rand.Float64() < math.Exp(-delta_d/temperature) {
			swap(i, j, perm)
			this_d += delta_d
			acceptance += 1
			if this_d < best_d {
				best_d = this_d
				best_p = perm
			}
		}

		// report progress
		if iter%period == 0 {
			fmt.Printf("%6d: temperature %v, acceptance %d best dist %v\n", iter, temperature, acceptance, best_d)
			temperature *= cooling
			acceptance = 0
		}
	}
	runtime := time.Since(start)
	distance := travelDist(best_p, dist)
	fmt.Printf("Found distance %v in %d steps, time %v\n", distance, acceptance, runtime)

	return best_p, distance
}
