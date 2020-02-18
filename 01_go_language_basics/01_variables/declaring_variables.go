// Declaring variables respects strong typing used within
// the language. Variables are defined within scopes to allow
// sharing or encapsulation of the values assigned to them.
// Variables can be set to reserve memory for a future value to come
// or can be directly initated with the respective values.
// Variables once set can not change their type, if you want to ensure
// that you also keep a given value constant you want to define a constant.

package main

import display "github.com/skillsmart/go/display"

func main() {
	//	To define a variable that can be accessed across all functions
	//	in a package, we define it here.. in {package scope}
	//  mynumber is defined to hold integer values
	var mynumber int
	display.ShowVariable(mynumber, "Initialize int variable without value")
	// It is created with a "null" value by go


	// We can "initiate" it to a float of 42.45
	var myfloat float32 = 42.45
	display.ShowVariable(myfloat, "Initialize float var with value 42.45")

	// Initializing a variable can also automatically
	// assign the correct type
	var mybool = true
	display.ShowVariable(mybool, "Initalize untyped var with 'true' value")

	//------------ Reassignment lock ----
	// Go helps you maintain correct reasoning over the
	// behavior and type of your variables. As such strong typing
	// will prevent you from reassigning types on a given variable.

	// This DOESN'T work... and you don't want it to eiter
	// (Trying to change boolean to integer)

	//  intbool := int(mybool)
	//  fmt.Prinln(intbool)


}

