package main

import (
	"fmt"
	"math/rand"
	"time"
)

const dataFile = "polygon.txt"

func main() {

	// variables
	var points [][2]float64
	var perm []int

	// read points
	//points = readFile(dataFile)
	points = makePolygon(10)
	npoints := len(points)

	// initialise distance matrix
	dist := make([][]float64, npoints)
	for i := range dist {
		dist[i] = make([]float64, npoints)
	}
	// compute distance matrix
	for i := 0; i < npoints; i++ {
		for j := 0; j < npoints; j++ {
			dist[i][j] = distance(points[i], points[j])
		}
	}
	// print distance matrix
	//for _, d := range dist {
	//	fmt.Println(d)
	//}

	// random permutation
	perm = rand.Perm(npoints)
	// perm = []int{0, 3, 6, 2, 5, 1, 4, 7}

	niters := int(1e06)
	start := time.Now()
	for iter := 0; iter < niters; iter++ {

		//old_d := travelDist(perm, dist)
		//fmt.Println(perm, old_d)
		i := rand.Intn(npoints)
		j := rand.Intn(npoints)
		//delta_d := deltaDist(i, j, perm, dist)
		//deltaDist(i, j, perm, dist)
		swap(i, j, perm)
		//new_d := travelDist(perm, dist)
		travelDist(perm, dist)
		// print errors
		//if math.Abs(old_d+delta_d-new_d) > 1e-10 {
		//	fmt.Println(perm, old_d, i, j, delta_d, new_d)
		//}
	}
	runtime := time.Since(start)
	fmt.Println(runtime)
}
