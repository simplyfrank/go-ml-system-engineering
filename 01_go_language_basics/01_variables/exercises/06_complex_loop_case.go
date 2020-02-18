// Print all numbers from 0 to 100
// Where you substitute:
// 1. All multiples of 3 with "Fizz"
// 2. All multiples of 5 with "Buzz"
// 3. And multiples of 3 & 5 with "FizzBuzz"

package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i %3 == 0 && i % 5 == 0:
			fmt.Println("FizzBuzz")
		case i % 3 == 0: fmt.Println("Fizz")
		case i % 5 == 0: fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
