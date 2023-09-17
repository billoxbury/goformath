package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func checkDeltaComp(perm []int, dist [][]float64, niters int, tolerance float64) {

	errCount := 0
	npoints := len(dist)
	for iter := 0; iter < niters; iter++ {

		old_d := travelDist(perm, dist)
		i := rand.Intn(npoints)
		j := rand.Intn(npoints)
		delta_d := deltaDist(i, j, perm, dist)
		swap(i, j, perm)
		new_d := travelDist(perm, dist)
		// print errors
		if math.Abs(old_d+delta_d-new_d) > tolerance {
			fmt.Println(perm, old_d, i, j, delta_d, new_d)
			errCount++
		}
	}
	fmt.Printf("Found %d errors out of %d at tolerance %v\n", errCount, niters, tolerance)
}

func timeTravelComp(perm []int, dist [][]float64, niters int) {

	npoints := len(dist)
	start := time.Now()
	for iter := 0; iter < niters; iter++ {
		i := rand.Intn(npoints)
		j := rand.Intn(npoints)
		swap(i, j, perm)
		travelDist(perm, dist)
	}
	runtime := time.Since(start)
	fmt.Printf("%d distance comps in time %v\n", niters, runtime)
}

func timeDeltaComp(perm []int, dist [][]float64, niters int) {

	npoints := len(dist)
	start := time.Now()
	for iter := 0; iter < niters; iter++ {
		i := rand.Intn(npoints)
		j := rand.Intn(npoints)
		deltaDist(i, j, perm, dist)
		swap(i, j, perm)
	}
	runtime := time.Since(start)
	fmt.Printf("%d delta comps in time %v\n", niters, runtime)
}
