package main

func fib(arr []int) {

	n := len(arr)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
}
