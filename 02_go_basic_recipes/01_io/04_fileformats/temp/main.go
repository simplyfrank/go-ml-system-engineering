package main

import (
	"encoding/gob"
	"fmt"
	temp "github.com/skillsmart/go/02_go_basic_recipes/01_io/04_fileformats/temp/temp"
	"io"
	"os"
)

func main() {

	// You write code to generate some values
	// and need to store in order to save on memory

	// You create yourself a temporary file in a new temporary directory
	tempDir, tempFile, err := temp.CreateTempFile("")
	fmt.Printf("Call to CreateTempFile resulting in: \nTemporary Directory location at: %v\nTemporary FileName: %s", tempDir, tempFile.Name())
	if err != nil {
		fmt.Println(err)
	}
	// You need to defer close to ensure proper handling
	defer os.RemoveAll(tempDir) // Using the filename to the temp dir, clear all generated files

	// Now we can use the temp file to write to and
	nums := []int{1,2,3,4,5,6,7,8,9}
	for i := range nums {
		// Let's use the square addition as an example for a mapper
		// where we want to ensure we keep intermediary results
		nums[i] += nums[i] ^ 2
		// In order to store intermediary results of your actions
		// we can now use the temp file

	}
	fmt.Printf("Numbers before encoding: %v", nums)
	// Encode to binary string
	encoder := gob.NewEncoder(tempFile)
	// Write to the file encoding the []int to binary
	if err := encoder.Encode(nums); err != nil && err != io.EOF{
		panic(err)
	}
	// In case of long running processes we could also store the tempfile in a
	// persistent store as key to a state of the array at this backup step
	// Here we are just reading it back out for checks

	decoder := gob.NewDecoder(tempFile)

	var p []int
	if err := decoder.Decode(&p); err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println(p)

	//decoder := gob.NewDecoder()
	//var storedArray []int
	//if err := decoder.Decode(&storedArray); err != nil {
	//	panic(err)
	//}

	//fmt.Printf("\n%v", storedArray)

}
	// Now that we have successfully transformed all elements we can quit
	// and the deferred function will clean up after us