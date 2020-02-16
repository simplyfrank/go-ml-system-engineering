// Besides conditional routing of control flows using if-else-then type
// statements, we can use switch statements to reduce control boiler plates

package main

import "fmt"

// Set which examples to run
var base bool = true
var checkDivisbility = false

func baseSwitch() {
	// We can set the variable externally
	states := []string{"running", "stopped", "pausing", "restarting"}

	desc := `Controling the flow of program execution based on the value of 'relevantState' that can 
can take a set of predefined values ['running', 'stopped', 'pausing', 'restarting']`
	fmt.Println(desc, "\n-------------------------------------------")

	for _, state := range states {
		fmt.Println("External relevantState has value: \t", state)

		// We want to determine the flow of the program based on the state
		switch state {
		// Now we use 'case' statement to check for specific states
		case "running":
			{
				fmt.Println("Status 'Running': Taking action X")
			}
		case "stopped":
			{
				fmt.Println("Status 'Stopped': Taking action Y")
			}
		case "pausing":
			{
				fmt.Println("Status 'Pausing': Taking action Z")
			}
		case "restarting":
			{
				fmt.Println("Status 'Restarting': Taking action X")
			}
		default:
			fmt.Println("Entered an out of list value.")
		}
		fmt.Println("-------------------------------")
	}t

}

func checkDivisibilitySwitch() {

	cnt2 := 0 // Initialize a new external variable to reference
	for cnt2 < 20 {
		// Let's check what the divisibility of the given number
		fmt.Printf("\nTesting %d: \n", cnt2)
		switch {
		case cnt2%2 == 0:
			fmt.Println("It's an even number")
		case cnt2%3 == 0:
			fmt.Println("It's an odd number")
		case cnt2%4 == 0:
			fmt.Println("Divisible by 4")
		case cnt2%5 == 0:
			fmt.Println("Divisible by 5")
		case cnt2%6 == 0:
			fmt.Println("Divisible by 6")
		case cnt2%7 == 0:
			fmt.Println("Divisible by 7")
		case cnt2%8 == 0:
			fmt.Println("Divisible by 8")
		case cnt2%9 == 0:
			fmt.Println("Divisible by 9")
		default:
			fmt.Println("It's a PRIME")
		}
		// And we still need to define the way we increment in the loop
		cnt2++
	}
}

func main() {

	// Basic syntax for the switch statement
	// The statement switches on the value of a given variable

	// Using a switch to replace a repetitive branching of if
	// statements on the condition of a given variable.
	if base == true {
		baseSwitch()
	}

	// Use switch statements to reduce boiler plate on parallel ifs
	// if cnt2 % 2 == 0 {
	//	  fmt.Println("It's an odd number")
	//} else if cnt2 % 3 == 0 { ....
	if checkDivisbility == true {
		checkDivisibilitySwitch()
	}
}
