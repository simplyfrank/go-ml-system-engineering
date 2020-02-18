package display

import "fmt"

// Show generates a format string to display the value shift
// in a set variable, the resulting type and a concise instruction
// describing the idea behind the exercise.
func ShowVariable(v interface{}, desc string) {
	fmt.Println(desc)
	fmt.Println("Variable has value: \t", v)
	fmt.Printf("variable has type: \t%T\n", v)
	fmt.Println("--------------\n")
}

// Iterate over all elements in cons and describe their
// type
func DescribeUntyped(cons ...interface{}) {
	fmt.Println("DISPLAY untyped variables | Descrition following iteratively\n\n")
	for i, val := range cons {
		fmt.Printf("Variable %d: \n", i)
		fmt.Println("-----------------------------")
		fmt.Printf("type: \t\t%T\nvalue:\t\t%v\n\n", val, val)
	}
}
