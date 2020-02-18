// Interfaces define abstract types by defining what functions
// need to be implemented on a value that is accepted to a function.
// This ensures that a clearly defined API can be enforced and can be
// programmed against.
//
// You can define a new interface from a given list of existing interfaces. Then
// all functions that are required by the interfaces need to be implemented in order
// to satisfy the composite interface.
//
// Interfaces are:
// 1. A set of methods
// 2. And the associated type
//
// Interfaces allow us to:
// 1. Ensure certain conditions are met on the input values
// 2. Certain behaviors have to be anticipated from a Go element
//
// ""
//"If you find yourself defining and using an interface in the same file, you
// are likely using interfaces the wrong way. """


package main

import (
	"fmt"
	"os"
	"strconv"
)
// To satisfy the io.Writer Interface
type Writer interface {
	Write(p []byte) (n int, err string)
}
// We will need to implement the function defined here 'Write()'
// on a custom Type EmptyWriter we want to define for our purpose.
type EmptyWriter os.File
func (w EmptyWriter) Write(p []byte) (n int, err string) {
	return 12, err
}

// To define a custom interface
// we follow the [type {NAME} interface]
// followed by the list of methods required
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Defining both a rectangle and a triangle as basic geometric
// types on which we can implement an area and perimeter function
type Rectangle struct {
	length float64
	height float64
	width float64
}
type Triangle struct {
	base float64
	height float64
}

// Implement Shape interface on the Rectangle by implementing the
// required methods
func (r Rectangle) Area() float64 {
	// This definition follows the required interface for Area()
	return 23.45 // --> return float64
}
func (r Rectangle) Perimeter() float64 {
	// This definition satisfies Perimeter() and fulfills the interface required for Shape
	return 24.355
}
func main() {

	// We can check if a given value satisfies a interface
	var myInt interface{} = 123
	fmt.Printf("The variable has type: %T\n with value: %d",  myInt, myInt)

	// See if the empty interface instantiated to int satisfies the int interface
	if k, ok := myInt.(int); ok {
		fmt.Println("Check if value satisfies interfaces:\n--------------------\n")
		fmt.Printf("Satisfies interface:?? -> %s\nWith value: \t%d\n\n", strconv.FormatBool(ok), k)
	}

	// Let's see if the interface for floats can also be satisfied
	fmt.Println("\nCheck if value satisfies interfaces for Float64:\n--------------------")
	if v, ok := myInt.(float64); ok {
		fmt.Println(v)
	} else {
		fmt.Println("Failed test without panicking.")
	}

	// Let us see if we satisfied the "Shape" Interface on the Rectangle Object
	var myRect interface{} = Rectangle {
		length: 24,
		height: 34,
		width:  12,
	}
	fmt.Println("\n>>> Check for successful implementation of Shape Interface:\n------------------------------------")
	if _, ok := myRect.(Shape); ok {
		fmt.Println("We successfully implemented the Shape Interface on Rectangle!")
	} else {
		fmt.Println("We are missing something here....")
	}
}

