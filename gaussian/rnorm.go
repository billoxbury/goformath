package main

// specify type: func that outputs a float
type genFloat func() float64

// function that takes a sample size and float generator and
// returns a sample serially
func rnorm(n int, rng genFloat) []float64 {

	var arr []float64
	for i := 0; i < n; i++ {
		arr = append(arr, rng())
	}
	return arr
}

// parallelised version using goroutines
func rnormPar(n int, w int, rng genFloat) []float64 {

	// sample size for each worker chunk
	// (keep things simple and assume it divides n exactly)
	m := n / w

	// intermediate slices and channel to return them
	var arr []float64
	rands := make(chan []float64, w)

	// split task into 'worker' goroutines and collate results
	for i := 0; i < w; i++ {
		go func() {
			rands <- rnorm(m, rng)
		}()
		x := <-rands
		arr = append(arr, x...)
	}

	// output results
	return arr
}
