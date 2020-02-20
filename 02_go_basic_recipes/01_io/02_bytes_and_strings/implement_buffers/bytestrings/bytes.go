package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
)

func WorkWithBuffer() error {
	fmt.Println("Results from Working with Buffer\n----------------------------------------")

	rawString := "easy to encode unicode into a byte array"
	fmt.Printf("rawString set to string: '%s'\n\n",rawString)

	b := Buffer(rawString)
	fmt.Printf("----->\nTransforming the string to type Buffer returns:\n \t\t\t\t%v\nWith Type:\t\t%T\n", b, b)

	// We can quickly convert a buffer back to bytes
	fmt.Printf("\n----->\nResult of byte transformation of the string:\n%v\n\n", b.Bytes())

	// Or a string
	fmt.Printf("---->\nReading buffer back to string results in:\n'%v'\nWith type: %T", b.String(), b.String())

	// as an io Reader we can make use of generic reader
	// functionality
	s, err := toString(b)
	if err != nil {
		return err
	}
	log.Println(s)

	// We can take the bytes and create a bytes reader passing it
	// the byte array converted rawString
	reader := bytes.NewReader([]byte(rawString))

	// We can plug it into a scanner for buffered reading and tokenization
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// then iterate over all scan events

	fmt.Println("\n\n---->>\nIterating over the returned buffer using ScanWords returns:")
	for scanner.Scan() {
		fmt.Printf("New scan returning: %s\n", scanner.Text())
	}

	return nil
}
