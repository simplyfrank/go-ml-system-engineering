package main

import (
	"fmt"
	"reflect"
)

func main() {
	// A way in order to iterate over the values in a struct
	// is to use the 'reflect' package to
	x := struct {
		Foo string
		Bar int
	}{"foo", 2}

	v := reflect.ValueOf(x)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	fmt.Println(values)
	fmt.Println(v)
	fmt.Printf(`
	Number of fields: %d
	
`, v.NumField(), v.)
}

