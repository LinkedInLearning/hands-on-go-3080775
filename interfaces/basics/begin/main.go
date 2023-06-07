// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak()
	walk()
}

// define a person type that implements humanoid interface
type person struct {
	name string
}

func (p person) speak() {
	fmt.Printf("I'm a person %s and I'm speaking\n", p.name)
}

func (p person) walk() {
	fmt.Printf("I'm a person %s and I'm walking\n", p.name)
}

// implement the Stringer interface for the person type
func (p person) String() string {
	return fmt.Sprintf("I'm a person with name %s\n", p.name)
}

// define a dog type that can walk but not speak
type dog struct {
}

func (d dog) walk() {
	fmt.Printf("I'm a dog and I can walk\n")
}

func main() {
	// invoke with a person
	p := person{name: "John Doe"}
	doHumanThings(p)

	// can we invoke with a dog?
	//doHumanThings(dog{})

	fmt.Println(p)
}

func doHumanThings(h humanoid) {
	h.speak()
	h.walk()
}
