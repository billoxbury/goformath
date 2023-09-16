package main

func distance(p1 [2]float64, p2 [2]float64) float64 {

	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

func travelDist(perm []int, dist [][]float64) float64 {

	td := 0.0
	np := len(perm)
	for i := range perm {
		td += dist[perm[i%np]][perm[(i+1)%np]]
	}
	return td
}
