package main

import (
	"math/rand"
)

const dataFile = "gb_cities.csv"

func main() {

	// variables
	var perm []int

	// read data
	labels, points := readCsv(dataFile)
	npoints := len(points)

	// ... or make data
	//npoints := 30
	//labels, points := makePolygon(npoints)

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
	// initialise random permutation
	perm = rand.Perm(npoints)
	// perm = []int{0, 3, 6, 2, 5, 1, 4, 7}

	niters := int(1e05)

	//checkDeltaComp(perm, dist, niters, float64(1e-10))
	//timeTravelComp(perm, dist, niters)
	//timeDeltaComp(perm, dist, niters)

	perm, _ = naiveSearch(perm, dist, niters)

	printRoute(perm, labels)
	writePerm(perm, "route.txt")
}
