package main

import "fmt"

// Idiomatic way to create Structs in GO
type person struct {
	name string
	age int
	member bool
}
func NewPerson(name string, age int, member bool) *person {
	p := person{
		name:   name,
		age:    age,
		member: member,
	}
	return &p
}

func main() {
	frank := NewPerson(
		 "Frank",
		34,
		false)
	claudia := NewPerson(
		"Claudia", 33, true)
	tobias := NewPerson(
		"Tobias", 39, false)

	fmt.Println(frank)
	fmt.Println(&frank)
	fmt.Println(claudia)
	fmt.Println(&claudia)
	fmt.Println(tobias, tobias.name, tobias.age)
	fmt.Println(&tobias.age, &tobias.name, &tobias.member)

}

func intSeq() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}


