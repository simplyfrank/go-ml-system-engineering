package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {

	// ---->>
	//Create a directory
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		fmt.Println(err)
	}

	// Create a file object that implements multiple interfaces
	// can be used as a reader or write if the correct bits are set
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Print(err)
	}

	// We can now wrie a known-length value to the file
	// and validate its correct execution
	value := []byte("hello")
	count, err := f.Write(value)

	if err != nil {
		fmt.Println(err)
	}

	if err := f.Close(); err != nil {
		fmt.Println(err)
	}

	// --->>
	//Validate all bytes have been written successfully
	if count != len(value) {
		fmt.Println(errors.New("incorrect length returned from write")
	}

	// ------>>
	//Reading from a file at once
	f, err = os.Open("testfile.txt"); if err != nil {
		fmt.Println(err)
	}
	// Stream file content to stdout
	io.Copy(os.Stdout, f)
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}



	// ------>
	// Change directories
	if err := os.Chdir(".."); err != nil {
		fmt.Println(err)
	}

	// ------->>
	// Perform cleanup. BE CAREFUL if you run at root
	if err := os.RemoveAll("example_dir"); err != nil {
		fmt.Println(err)
	}
}