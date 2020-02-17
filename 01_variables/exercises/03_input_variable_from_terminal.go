// Exercise:
// Request input from user from the command line to define his/her name
// and print a customized greeting back to them.

// My solution......
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please enter your name:")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("\nWelcome to the show %s\n", input)
}

