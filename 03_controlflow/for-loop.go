package main

import (
	"fmt"
	//display "github.com/skillsmart/go/display"
)

func main() {

	// Go uses a single generic way to define looping constructs.
	// This construct defines three attributes
	// 1. The variable instantiation ( i:= ??)
	// 2. The continue clause ( i <|>|<=|>=) to specify when to
	// continue looping
	// 3. The way incrementer to set the value at each iteration of the
	// loop

	// Loop for 10 times
	fmt.Println("This loop will be running for 10 Iterations\n-------------------------------\n")
	for i := 0; i < 11; i++ {
		fmt.Println("This is loop number ", i)
	}

	// If we don't specify the loop variable, the break clause and the incrementer function
	// we end up defining a "while function"


	fmt.Println("\nThis loop will count from 0 to 100 only reporting on each 25th event and break at 100")
	fmt.Println("----------------------------------------\n")
	// To alter an external variable
	cnt := 0
	// This construct will loop forever
	for {
		cnt++
		if cnt%25 == 0 {
			fmt.Println("Reaching a count divisble by 25: ", cnt)
		}
		if cnt >= 100 {
			println("Reached 100! I am breaking the loop.")
			break
		}
	}

	// We can move the break statement up into the loop declaration
	// and still reference a variable outside of the loop scope



}
