// ATTENTION: This is an advanced topic, please skip when you are just getting started
// as it breaks initially helpful abstractions that allow you to get started easily!

// Content:
// Constructs are objects with "arbitrary types" and as such do not easily
// conform with the strict type safe interface of go. If you want to manipulate
// them in programmatic ways, such as ranging over them, using them for key/value bases
// lookups and altering their runtime behavior you need to look towards "reflection" for a solution.

// Definition of reflection:
// Read the GoBlog https://blog.golang.org/laws-of-reflection for futher details
//
package main

// Metaprogramming enables a program to examine its own structure.
// In go this strongly builds on the inspection of types and their attriibutes
// and relationships.

// Go is statically typed, and as such the type of all variables in the program
// are defined and known at compile time.
// Even the use of the same underlying types does not allow us to
// use different types interchangeably

// Remember: This doesn't work
//	i := j // Cannot use 'j' (type myInt) as type int in assignment

// Regarding structs we face the same underlying challenge
// If we define structs, we are declaring a variety of types under
// a compound structure
