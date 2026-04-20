package main

import (
	"fmt"
)

//y:=11 this doesn't work coz short hand decleration only works inside function.

var x int = 101

//fmt.Println("learning all about variables in go", x) this does not work coz, In Go programming language,
//  a .go source file at the top level (package scope) can contain declarations, not executable statements.

func main() {
	fmt.Println("Printing value of x", x)
	x := 100
	fmt.Println("value of x", x)
	explicit_type_and_value()
	type_inferred()
	zero_value()
	shorthand_declaration()
}

func explicit_type_and_value() {
	var name string = "Siddhartha Reddy"
	fmt.Print(name)
}

func type_inferred() {
	var name = "Siddhartha Reddy"
	fmt.Printf("%s", name)
}

func zero_value() {
	var name string
	var num int
	var decimal float64
	var ok bool
	fmt.Println(name, num, decimal, ok)
	fmt.Println("**********************")
	var n int
	var s string
	var b bool
	var nums []int
	fmt.Printf("n=%d, s=%q, b=%v, nums=%v, nums==nil:%v\n", n, s, b, nums, nums == nil)
}

func shorthand_declaration() {
	name := "Sid"
	phone := 1234567890
	fmt.Println(name, phone)
}

// If you have observed that I deliberately did not call this function in the main function coz.
// In Go, unused local variables and unused imports cause compile errors.
// Unused functions, types, constants, and package-level variables are usually allowed.
func dec_multiple_variable() {
	var a, b, c, d, e int = 1, 2, 3, 4, 5
	c, d = d, c
	fmt.Println(a, b, c, d, e)
}

// _ is the blank identifier in Go, used to ignore values you do not need.
// Commonly used with functions returning multiple values: _, err := someFunc().
// Often used in loops: for _, v := range items { } to ignore index.
// Go does not allow unused variables, so _ helps discard unwanted values.
// _ can receive a value, but it cannot be read or printed.
