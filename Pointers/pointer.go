package main

import "fmt"

func main() {
	fmt.Println("-----Learning about pointers-------")
	declaration()
	update_value()
	nil_pointer()
	update_value_by_function()
	pointer_struct_demo()
	ting()
	array()
	slice()
	pointer_receiver_interface()
}

func declaration() {
	var x int = 10

	var p *int = &x
	// *int means “pointer to int” type
	// p can store the memory address of an int variable

	// &x means “address of x”
	// it returns the memory address where x is stored

	// * before a TYPE means pointer type declaration
	// * before a VARIABLE means dereference

	fmt.Println("Address of x is stored in p and the value of x and p is:", x, ",", p)
}

func update_value() {
	x := 10

	fmt.Println("x:", x)
	fmt.Println("address of x:", &x)

	p := &x

	fmt.Println("p stores:", p)
	fmt.Println("value pointed by p:", *p)

	*p = *p + 20

	fmt.Println("updated x:", x)
	fmt.Println("updated dereferenced value:", *p)
}

func nil_pointer() {
	var p *int

	fmt.Println(p) // prints <nil>

	// fmt.Println(*p) results in error: panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x10217f840]
}

func update(p *int) {
	*p = 100

	fmt.Println("Println value of p in update function", p)
}

func update_value_by_function() {
	x := 10

	p := &x

	fmt.Println("value of x before passing to update function:", x)
	fmt.Println("Printing value of p before update() function:", p)

	update(p)

	fmt.Println("value of x after passing to update() function", x)
}

type user struct {
	name string
	age  int
}

// Pointer receiver
func (u *user) call() {

	// modifies original struct
	u.age = 22
	u.name = "reddy"
}

func pointer_struct_demo() {

	// value type
	u := user{"Sid", 12}

	// Go automatically converts:
	// u.call() -> (&u).call()
	u.call()

	p := &u

	p.call()

	// user{"damu", 23}.call() doesn't work.

	fmt.Println(u)
}

func ting() {
	x := 10

	p := &x

	pp := &p

	fmt.Println("multiple dereferencing:", **pp)
}

func array_pointer(a *[4]int) {
	a[0] = 1
}

func array() {
	arr := [4]int{2, 3, 4, 5}

	fmt.Println("array:", arr)

	array_pointer(&arr)

	fmt.Println("after passing to function:", arr)
}

func slice_pointer(a []int) {
	a[0] = 1
}

func slice() {
	arr := []int{2, 3, 4, 5}

	fmt.Println("slice:", arr)

	slice_pointer(arr)

	fmt.Println("after passing to function:", arr)
}

//Advance, learn interface and then do the below excercise.

type Speaker interface {
	speak()
}

type hitler struct{}

// Pointer receiver
func (h *hitler) speak() {
	fmt.Println("hitler speech")
}

func pointer_receiver_interface() {

	// ❌ This gives compile error because
	// methods with pointer receivers belong
	// only to the method set of *T, not T.

	// var s Speaker = hitler{}

	// ✅ Works because *hitler implements Speaker
	var s Speaker = &hitler{}

	s.speak()

	// If speak() had VALUE receiver:
	//
	// func (h hitler) speak()
	//
	// then BOTH would work:
	//
	// var s Speaker = hitler{}
	// var s Speaker = &hitler{}
	//
	// because value receiver methods belong
	// to the method sets of both T and *T.
}
