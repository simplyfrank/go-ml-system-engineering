package main

// A custom type can implicitly implement a varying number
// of interfaces. This is done by associating a number of
// functions with a custom type

// Let's create a custom Sequence type to manage a slice of ints
type Sequence []int

// Let's implement the sortable Interface
// Satisfying the list of required functions and the accepted
// attributes, their types and the type of returns automatically
// allows us to use our custom type with all functions that expect the
// 'sortable' interface
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i,j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}



func main() {

}
