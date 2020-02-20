package main

import (
	"bytes"
	"fmt"

	interfaces "github.com/SkillSmart/go/02_go_basic_recipes/01_io/interfaces"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())

	fmt.Print("stdout on PipeExample =")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
