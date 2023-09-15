package main

// a dynamic programming implementation
func fib(arr []int) {

	n := len(arr)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
}

// a naive and fairly disastrous recursive implementation
func fibNaive(n int) int {

	if n <= 1 {
		return n
	} else {
		return fibNaive(n-1) + fibNaive(n-2)
	}
}
