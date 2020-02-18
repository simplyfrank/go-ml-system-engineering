// Exercise:
// Ask your user for a divisor and divident and calculate the rest left from
// the division and return it back to the user.
// Ensure that the program handles wrong/missing inputs to the program


// My Solution:

package main

import (
	"fmt"
	"os"
)

func main() {
	var divisor int
	var divident int
	fmt.Println("Input divisor: ")
	if _, err := fmt.Scanln(&divisor); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("Input divident: ")
	if _, err := fmt.Scanln(&divident); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// Calculate and return
	if divident > 0 {
		rest := divisor % divident
		fmt.Printf("Your division leaves a rest of: %d", rest)
	} else {
		fmt.Println("Division is undefined for the entered values. Restart!")
	}
}
