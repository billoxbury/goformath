package main

import (
	"math"
)

// compute L_2 distance
func distance(p1 [2]float64, p2 [2]float64) float64 {
	return math.Sqrt((p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1]))
}

// total distance around a given route
func travelDist(perm []int, dist [][]float64) float64 {

	td := 0.0
	np := len(perm)
	for i := range perm {
		td += dist[perm[i%np]][perm[(i+1)%np]]
	}
	return td
}
