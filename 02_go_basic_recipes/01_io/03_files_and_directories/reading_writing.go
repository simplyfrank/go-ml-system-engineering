package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// Paths are always resolved based on your calling environment
	// i am using an IDE here that calls each file from the project root
	// Adjust for your environment if the file is not found
	data, err := ioutil.ReadFile("./data/string_data.txt")
	if err != nil { panic(err) }
	fmt.Println(string(data))

	// ------>
	//Write to a file
	teststring := "This is my teststring"
	err = ioutil.WriteFile("data/hello.dat", []byte(teststring), 0644)
	if err != nil { panic(err) }

	// Reading it back out for testing
	data, err = ioutil.ReadFile("data/hello.dat")
	if err != nil { panic(err) }
	fmt.Println("Read out: ", string(data))



	// Using a tempfile for quick writing and reading
	tf, err := ioutil.TempFile("data/tmp", "tmp")
	if err != nil { panic (err) }
	defer tf.Close()

	// Writing a temp message to the file
	fmt.Println("Writing message to temp file...\n")
	tf.Write([]byte("Hey there this is a temporary message. I will soon not need it any more"))

	// Reading the temp message back out
	tf_data, err := ioutil.ReadFile(tf.Name())
	if err != nil { panic(err) }

	fmt.Println("Reading back out the temp file: \n", string(tf_data))
}