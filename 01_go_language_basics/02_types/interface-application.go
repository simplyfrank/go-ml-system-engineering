// We are discussing how to build up a more complex method abstraction interface
// over the analogy of modeling a shepherds need to train either a dog/wolf as a SheepDog
// abstracting away differences between the animals by just focussing on the expected behavior
// that every SheepDog has to exhibit.

// Content:
// 1. Structs as types in go, and how this impacts assignability
// 2. Interfaces as a way to abstract over types by enforcing common behavior
// 3. Practical example of a function definition build on top of an abstracted interface


package main


type Animal struct {
	gender string
	weight float32
	height float32
	family string
	registered bool
}

// And even when we define a new struct with the same underlying
// attributes, the two can not be used interchangeably
type Pet struct {
	gender string
	weight float32
	height float32
	family string
	registered bool
}

// Interfaces are designed to allow the abstract use of types
// along the lines of a fixed, defined method interface, that allows
// to build on the idea of 'duck-typing' and treat types as equal as long as
// they 'exhibit' the same looks and behaviour.

// Regarding both wild animals and pets as partners for work and we want to
// ensure that we can tell them to sit, go somewhere, fetch something and we can
// ask them about their mood which they can signal us, we assume to not care about
// wether they are wild or domesticated.

// 1. We define the expected interface
type SheepDog interface {
	Sit() bool
	Move(target string) bool
	Fetch(target string) bool
	Mood() string
	Speak() string
	GatherHerd(radius float32) bool
}
// 2. All types that provide these methods are automatically set conformant to the interface
func (Animal) Sit() bool { return true}
func (Animal) Move(target string) bool { return true}
func (Animal) Fetch(target string) bool { return true}
func (Animal) Mood() string { return "Grrrr....."}
func (Animal) Speak() string { return "Wooof"}
func (Animal) GatherHerd(radius float32) bool { return true }

func (Pet) Sit() bool { return true}
func (Pet) Move(target string) bool { return true}
func (Pet) Fetch(target string) bool { return true}
func (Pet) Mood() string { return "Grrrr....."}
func (Pet) Speak() string { return "Wooof"}
func (Pet) GatherHerd(radius float32) bool { return true }

// 3. This will now allow us to define the structure of our program
// to be flexible against if we get a wolf or a dog for working the herd
func main() {

	type myInt int
	var i int = 12
	var j myInt = 45
	_,_ = i,j // The assignment to _ tells the compiler we will not use the variables (i,j)

	myWolf := Animal{
		gender:     "male",
		weight:     45,
		height:     50,
		family:     "dog",
		registered: false,
	}
	myPet := Pet{
		gender:     "male",
		weight:     45,
		height:     50,
		family:     "dog",
		registered: false,
	}
	// This doesn't work
	//	myAnimal = myPet //Cannot use 'myPet' (type Pet) as type Animal in assignment

	//Using the interface of "SheepDog" we are now independent of the type
	sheepWolf := myWolf
	sheepDog := myPet

	// We can use both of them to perform the same action
	sheepWolf.GatherHerd(24.5)
	sheepDog.GatherHerd(24.5)

	// Also both animals will satisfy the interface for the GatherClose function
	GatherClose(sheepDog)
	GatherClose(sheepWolf)


	// This also allows us to group them in data structures like an
	// array, slice, map, etc...
	for _, workDog := range []SheepDog{sheepDog, sheepWolf} {
		GatherClose(workDog)
	}


}

func GatherClose(worker SheepDog) bool {
	// We instruct the worker which has to be 'trained' as a SheepDog
	// to collect the herd iteratively til it is within a given target radius

	// We know we can "expect" a SheepDog to know how to gather a herd
	if ok := worker.GatherHerd(14); ok {
		return true
	} else {
		return false
	}
}

