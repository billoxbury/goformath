package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

// make polygon point set
func makePolygon(n int) ([]string, [][2]float64) {

	var points [][2]float64
	var labels []string
	var x, y float64

	for i := 0; i < n; i++ {
		x = math.Cos(2.0 * float64(i) * math.Pi / float64(n))
		y = math.Sin(2.0 * float64(i) * math.Pi / float64(n))
		pt := [2]float64{x, y}
		labels = append(labels, strconv.Itoa(i))
		points = append(points, pt)
	}
	return labels, points
}

// read data file into points slice
func readCsv(dataFile string) ([]string, [][2]float64) {

	var points [][2]float64
	var labels []string
	var df *os.File

	df, _ = os.Open(dataFile)

	// create scanner (bufio)
	scanner := bufio.NewScanner(df)
	scanner.Scan() // this skips the first line
	for scanner.Scan() {

		record := strings.Split(scanner.Text(), ",")
		labels = append(labels, record[0])
		x, _ := strconv.ParseFloat(record[1], 64)
		y, _ := strconv.ParseFloat(record[2], 64)
		pt := [2]float64{x, y}
		points = append(points, pt)
	}
	df.Close()
	return labels, points
}
