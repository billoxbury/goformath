/*

Run with

go build -o tsp ./src

./tsp -poly 17
./tsp -poly 17 -naive // look for 2*Pi

./tsp -dat ./data/gb_cities.csv -niters 100000000 -v
Rscript ./R/drawRoute.R ./data/gb_cities.csv ./data/route.txt ./img/map.pdf


*/

package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {

	// variables
	var dataFile, outFile string
	var poly, niters int
	var npoints int = 0
	var points [][2]float64
	var labels []string
	var naive bool
	var verbose bool

	// cmd line arguments
	flag.StringVar(&dataFile, "dat", "", "cities file (CSV)")
	flag.StringVar(&outFile, "out", "route.txt", "output file")
	flag.IntVar(&poly, "poly", 0, "polygon size (option)")
	flag.IntVar(&niters, "niters", int(1e06), "max iterations for search")
	flag.BoolVar(&naive, "naive", false, "naive search")
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.Parse()

	// read or make data
	if dataFile != "" {
		labels, points = readCsv(dataFile)
		npoints = len(points)
	} else if poly > 0 {
		npoints = poly
		labels, points = makePolygon(npoints)
	}
	if npoints == 0 {
		fmt.Println("No problem to process")
		return
	}

	// initialise
	dist := distMatrix(points)
	perm := rand.Perm(npoints)

	// VARIOUS TESTS
	//checkDeltaComp(perm, dist, niters, float64(1e-10))
	//timeTravelComp(perm, dist, niters)
	//timeDeltaComp(perm, dist, niters)

	if naive {
		// NAIVE SEARCH
		perm, _ = naiveSearch(perm, dist, niters)
	} else {
		// METROPOLIS SEARCH
		perm, _ = metropolisSearch(perm, dist, niters, verbose)
	}

	// return results
	printRoute(perm, labels)
	writePerm(perm, "./data/"+outFile)
}
