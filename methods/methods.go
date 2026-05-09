package main

import (
	"fmt"
)

func main() {

	// Just calling all examples one by one
	fmt.Println("-----Methods------")

	methods_declaration()
	function_type()
	interface_methods()
}

// ----------------------------------------------------
// METHODS
// ----------------------------------------------------

// Simple struct
type sid struct {
	Name string
	Age  int
}

// Value receiver.
// Whatever changes we do here happen on a copy.
func (s sid) intro() {
	fmt.Println("My Name is:", s.Name)
}

// Again value receiver.
//
// r here is not the original object.
// Go copies the struct before calling the method.
func (r sid) update_age() {

	// Only changing local copy
	r.Age = 25

	fmt.Println("Inside method:", r)
}

// Pointer receiver.
//
// Here r points to original struct,
// so changes reflect outside too.
func (r *sid) update_age_pointer() {
	r.Age = 25
}

func methods_declaration() {

	// Named initialization
	s := sid{
		Name: "Siddhartha",
	}

	// Positional initialization
	r := sid{
		"Reddy",
		24,
	}

	// Partial initialization
	t := sid{
		Age: 26,
	}

	s.intro()

	// Since update_age() uses value receiver,
	// original r will stay unchanged.
	r.update_age()

	fmt.Println("After value receiver call:", r)

	// This method needs *sid,
	// but Go automatically takes address of r.
	//
	// Internally it becomes:
	// (&r).update_age_pointer()
	r.update_age_pointer()

	fmt.Println("After pointer receiver call:", r)

	// Explicit pointer style
	(&t).update_age_pointer()

	fmt.Println(t)
}

// ----------------------------------------------------
// FUNCTION TYPES
// ----------------------------------------------------

// handler is a custom function type.
//
// Any function matching this signature
// can be assigned to it.
type handler func(a int, b string)

// Methods can even be attached to function types.
func (h handler) ab(a int, b string) {

	// h itself is the function
	h(a, b)
}

// Normal function matching handler signature
func info(a int, b string) {
	fmt.Println("Hey dude", a, b)
}

func function_type() {

	// Anonymous function assigned to handler
	h := handler(func(a int, b string) {
		fmt.Println("Hey there", a, b)
	})

	// Normal function call
	h(10, "Sid")

	// Calling method attached to function type
	h.ab(11, "Sid")

	var f handler

	// Assigning regular function
	f = info

	f(12, "dis")

	f.ab(10, "reddy")
}

// ----------------------------------------------------
// INTERFACES + METHOD SETS
// ----------------------------------------------------

// Any type having what() method
// satisfies this interface.
type Name interface {
	what()
}

type s string
type n string

// Value receiver.
//
// Both s and *s can use this method.
func (s s) what() {
	fmt.Println(s)
}

// Pointer receiver.
//
// Only *n gets this method in method set.
func (n *n) what() {
	fmt.Println(*n)
}

func interface_methods() {

	var sid Name

	// s satisfies interface directly
	sid = s("siddhartha")

	sid.what()

	var red Name

	temp := s("reddy")

	// *s also satisfies interface.
	//
	// Pointer method set includes
	// value receiver methods too.
	red = &temp

	red.what()

	siva := n("siva")

	p := &siva

	// n itself does NOT implement Name.
	//
	// Because what() belongs only to *n.
	var si Name = p

	si.what()

	// This would fail:
	//
	// var x Name = siva
	//
	// because n does not implement Name.
}
