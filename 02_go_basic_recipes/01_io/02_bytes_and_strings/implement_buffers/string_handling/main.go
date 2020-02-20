package main

import (
	bytestrings "github.com/SkillSmart/go/02_go_basic_recipes/01_io/02_bytes_and_strings/implement_buffers/bytestrings"
)

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	// Each of these print to stdout
	bytestrings.SearchString("", "")
	bytestrings.ModifyString("")
	//bytestrings.StringReader()
}