// COMMAND LINE INTERFACE - Example using the flag package
//
//This recipe attempts to demonstrate most of the common usages of the flag package. It shows custom variable types, a
//variety of built-in variables, shorthand flags, and writing all flags to a common structure. This is the first recipe
//to require a main function, as the main usage of flag (flag.Parse()) should be called from main. As a result, the normal
//example directory is omitted.
//
//The example usage of this application shows that you get -h automatically to get a list of flags that are included.
//	Some other things to note are Boolean flags that are invoked without arguments, and that the flag order doesn't matter.
//The flag package is a quick way to structure input for command-line applications and provide a flexible means of
//specifying upfront user input for things such as setting up logger levels or the verbosity of an application. In the
//Using command-line arguments recipe, we'll explore flagsets and switch between them using arguments.


package main

import (
	"flag"
	"fmt"
)

func main() {
	// initialize our setup
	c := Config{}
	c.Setup()

	// generally call this from main
	flag.Parse()

	fmt.Println(c.GetMessage())
}
