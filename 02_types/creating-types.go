// Let's have a look at how we can define custom types in go.
// There are multiple 'idiomatic' best-practices that we can follow
// to ensure we do not simply 'shadow' types with alternative declarations

// Goals to achieve in defining custom types:
// 1. Provide compile time checks on the values associated with a given variable
// 2...
package main

import "fmt"

func main() {
	// Types can be derived from core types in go
	type positiveInt int
	type negativeInt int

	// Both custom types can be used where ever we can use
	// the derived custom type

	// When we define two custom variables
	var numA positiveInt = 12
	var numB negativeInt = -24

	// We can not apply it to the sum function that expects
	// to be called with a multiple numbers of integers
	//result := func (a,b int) int {
	//	return a + b
	//}(numA, numB)
	//fmt.Println(result)
	// (uncomment to check!)

	// But we can convert the values to int as they share a common
	// underlying type (from specific up to general type)
	var a int = int(numA)
	var b int = int(numB)

	// Using the variables after casting works
	res := func(a,b int) int {
		return a + b
	}(a,b)
	fmt.Println(res)

	// --> Remember: Custom types provide insurance that values are
	// checked appropriately during instantiation and we can rely on their
	// attributes when working with them within a function
}