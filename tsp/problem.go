package main

import (
	"fmt"
	"math"
	"os"
)

// make polygon point set
func makePolygon(n int) [][2]float64 {

	var points [][2]float64
	var x, y float64

	for i := 0; i < n; i++ {
		x = math.Cos(2.0 * float64(i) * math.Pi / float64(n))
		y = math.Sin(2.0 * float64(i) * math.Pi / float64(n))
		pt := [2]float64{x, y}
		points = append(points, pt)
	}
	return points
}

// read data file into points slice
/*
func normalPercentiles() []float64 {

	// read data
	var qf *os.File
	qf, _ = os.Open(quantileFile)
	var normalQ []float64

	// create scanner (bufio)
	scanner := bufio.NewScanner(qf)
	for scanner.Scan() {

		line := scanner.Text()
		x, _ := strconv.ParseFloat(line, 64)
		normalQ = append(normalQ, x)
	}
	qf.Close()

	return normalQ
}
*/

func readFile(dataFile string) [][2]float64 {

	var points [][2]float64
	var x, y float64
	var df *os.File

	df, _ = os.Open(dataFile)
	for {
		n, _ := fmt.Fscanln(df, &x, &y)
		if n == 0 {
			break
		}
		pt := [2]float64{x, y}
		points = append(points, pt)
	}
	df.Close()
	return points
}
