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

	dist := distMatrix(points)

	// initialise random permutation
	perm = rand.Perm(npoints)

	niters := int(1e08)

	// VARIOUS TESTS
	//checkDeltaComp(perm, dist, niters, float64(1e-10))
	//timeTravelComp(perm, dist, niters)
	//timeDeltaComp(perm, dist, niters)

	// NAIVE SEARCH
	//perm, _ = naiveSearch(perm, dist, niters)

	// METROPOLIS SEARCH
	perm, _ = metropolisSearch(perm, dist, niters)

	printRoute(perm, labels)
	writePerm(perm, "route.txt")
}
