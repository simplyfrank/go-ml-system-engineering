package main

import "fmt"

func main() {
	// A struct is implemented as a keyword in go.
	type Coordinate struct {
		// You can define values on the struct across all
		// available types
		lat float64
		lng float64
	}

	// The struct now controls the type of values stored on the
	// compound data type
	var homeLocation Coordinate = Coordinate{
		lat: 23.4,
		lng: 45.5,
	}
	// You can accesss the underlying values stored in the
	// struct using the {Var}.{Value} notation
	fmt.Println("Exercise: 1 --> Define Coordinate struct\n------------------------")
	fmt.Println("Home lattitude set to: ", homeLocation.lat)
	fmt.Println("Home longitude set to: ", homeLocation.lng)

	// Instead of rebuilding logical packages for your values for more complex structs,
	// we can also just nest them
	type Location struct {
		location Coordinate
		name string
		rating int
		visited bool
	}
	// When instantiating the position value, we are now asked to provide
	// a nested struct of type "Coordinate"
	var bestPizzaPlace Location = Location{
		location: Coordinate{ // Here we are instantiating a new Value of type Coordinate
			34.5,
			45.5,
		},
		name: "PizzaPlace",
		rating: 5,
		visited: true, // Go requires the use of a trailing semi-colon
	}
	// When printing the nested type we are getting return of all values stored on
	// the type. Nested values are displayed within curly braces
	fmt.Println(bestPizzaPlace)
	// We can acccess the nested values by repeating the dot lookup
	fmt.Println(bestPizzaPlace.location.lat)
	fmt.Println(bestPizzaPlace.location.lng)

	// structs are not represented directly with a map interface, and you can not
	// use them to easily range over all keys and values in the struct.

	//for k,v := range bestPizzaPlace {
	//	// This doesn't work
	//}



}
