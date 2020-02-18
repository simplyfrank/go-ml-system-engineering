// If we list all the natural numbers below 10 that are
// multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these
// multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.


package main

import (
	"fmt"
)

func main() {
	// Set variables to be used in the program
	var low, high int = 1, 1_000_000
	var moduli []int = []int{3,5}
	// Get the list of multiples matching all cases
	var multiples []int = getmultiples(low, high, moduli...)
	// Calculate the sum product over all resulting multiples
	prod := sum(multiples)
	// Display the result
	fmt.Printf("The product of all numbers between: \t%d and %d\n", low, high)
	fmt.Printf("Selected to all multiples of : \t\t%v\n", moduli)
	fmt.Printf("Number of matching numbers: \t\t%d\n", len(multiples))
	fmt.Printf("Matching numbers: \t\t%v\n", multiples)
	fmt.Printf("Resulting sum: \t\t%d\n", prod)
}


// Function accepts a varying list of integers and returns an array of int
// in the range of low to high
func getmultiples(low int, high int, nums ...int) []int {
	multiples := []int{}
	out:
	for i := low; i <= high; i++ {

		for _,j := range nums {
			// Check if the i is divisible by j
			if i % j != 0 {
				continue out
			}
		}
		// if i was divisible by all j add it to the list of multiples
		multiples = append(multiples, i)
		}
	return multiples
}

func sum(nums []int) int64 {
	var product int64 = 1
	for _, val := range nums {
		product += int64(val)
	}
	return product
}
