package main

import (
	display "github.com/skillsmart/go/display"
)

func main() {
	// In order to define variables that keep their value
	// as well as their type across the entire lifetime, go
	// provides you with "contants"

	// Instantiate a single const
	const myConstBool bool = true
	display.ShowVariable(myConstBool, "Boolean type constant initialized to true")

	// Again we can automaticall infer the type
	const myStringConst = "A saying that stays true with you."
	display.ShowVariable(myStringConst, "Automatically infered type from initialization")

	// Constants can be assinged to variables of their own type only
	// 1) This works
	var flexString string = myStringConst
	display.ShowVariable(flexString, "Initialized string var with string constant")

	// 2) This is not allowed (Uncomment to test out)
	//var flexInteger int = myStringConst


	// Constants can be defined as untyped constants, in which case
	// their default type will be inferred based on the initialization values
	const untypedInt = 23
	const untypedStr = "Untyped String"
	const untypedBool = true
	const untypedFloat = 4334.5234

	display.DescribeUntyped(
		untypedInt, untypedStr, untypedBool, untypedFloat)


}
