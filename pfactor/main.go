/*

- static arrays (can't be declared const; use of [...])
- math package
- cast int as float64
- slices and append()
- complex for loop "for a,b := X,Y; ... ; a,b = P,Q { }"
- infinite loop for {}

Run with:

> go build .
> ./pfactor

*/

package main

import (
	"fmt"
)

func main() {

	// interaction loop
	fmt.Println("Enter number to factor, or Ctrl-C to quit:")
	for {
		// next input
		fmt.Printf("pfactor> ")
		var n int
		fmt.Scanln(&n)

		// factorisation
		fmt.Println(find_all_factors(n))
	}
}
