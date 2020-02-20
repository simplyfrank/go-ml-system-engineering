package bytestrings

import (
	"fmt"
	"strings"
)

// Here are a number of ways to search strings
func SearchString(s string, target string) {
	// This sets s with a default argument
	if s == "" {
		s = "This is a test string"
	}

	// Search for a containing word
	var res bool
	res = strings.Contains(s, target)
	fmt.Printf("String: %s\nContains %s? %v", s, target, res)

	// Check if string contains any part of it
	res = strings.ContainsAny(s, "abc")
	fmt.Printf(`ContainsAny check:\nDoes %s\n contain any chars of %s? Result: %v`, s, "abc", res)

	// Check for prefix
	res = strings.HasPrefix(s, "This")
	fmt.Printf("HasPrefix checks:\nDoes %s\nstart with 'This'?\nResult: %v", s, res)

	// Check for suffix
	res = strings.HasSuffix(s, "ing")
	fmt.Printf("HasSuffix checks:\nDoes %s end with 'ing'?\nResult: %v", s, res)
}

func ModifyString(s string) {
	fmt.Println("Displays a series of string modifications\n-------------------------------\n\n")
	if s == "" {
		s = "simple testing string"
	}

	// Split string based on empty space
	splitString := strings.Split(s, " ")
	fmt.Printf(`Split splits a string into elements based on a seperator\n
Splitting the string "%s" on seperator " " results in: %v'\nWith type: %T`, s, splitString, splitString)
}
