package main

import (
	"bufio"
	"os"
	"strconv"
)

const quantileFile = "normalQ.txt"

func normalQuantiles() []float64 {

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
