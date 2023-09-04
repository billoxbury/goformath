package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

const quantileFile = "normalQ.txt"

// read from file the theoretical normal percentiles:
func normalPercentiles() []float64 {

	// goodness-of-fit tests
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

// compute mean:
func mean(x []float64) float64 {
	n := float64(len(x))
	mu := 0.0
	for i := 0; i < len(x); i++ {
		mu += x[i]
	}
	mu /= n
	return mu
}

// compute variance:
func variance(x []float64) float64 {

	n := float64(len(x))
	mu := mean(x)
	v := 0.0
	for i := 0; i < len(x); i++ {
		v += (x[i] - mu) * (x[i] - mu)
	}
	v /= n
	return v
}

// compute percentiles of a sample:
func percentiles(x []float64) []float64 {

	var pc []float64
	n := len(x)

	sort.Float64s(x)
	for i := 1; i < 100; i++ {
		pc = append(pc, x[i*n/100])
	}
	return pc
}

// compare actual with theoretical percentiles and return
// 'deviance' = sum of square differences
func deviance(x []float64) float64 {

	pc0 := normalPercentiles()
	pc1 := percentiles(x)
	dev := 0.0
	for i := 0; i < 99; i++ {
		dev += (pc0[i] - pc1[i]) * (pc0[i] - pc1[i])
	}
	return dev
}
