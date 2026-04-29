package main

import "fmt"

func main() {
	fmt.Println("-------Learning Arrays----------")
	declaration()
	access_update_elements()
	compare()
	iteration()
	modify()
	memory_layout()
}

func declaration() {
	var arr [3]int = [3]int{}
	fmt.Println("printing arr array", arr)

	var a [2]string
	a = [2]string{"sid", "reddy"}
	fmt.Println("printing arrray a", a)

	b := [...]float32{1.23, 4.24, 3}
	fmt.Println("printing array b", b)
}

func access_update_elements() {
	// Declare an array of 3 elements of type int8
	// At this point, all values are initialized to zero → [0 0 0]
	var a [3]int8

	// Assign values to the array
	// Note: int8 range is -128 to 127, so -128 is valid
	a = [3]int8{1, 2, -128}

	// Accessing element using index
	// Arrays are zero-indexed → a[0] is the first element
	fmt.Println("1st element in a array:", a[0])

	// Print the memory address of array 'a'
	// This is the address of the entire array (not individual elements)
	fmt.Printf("address of a: %p \n", &a)

	// Update element at index 2
	// Now array becomes → [1 2 3]
	a[2] = 3

	// Assign 'a' to 'b'
	// IMPORTANT: Arrays use value semantics → this creates a full copy
	// 'b' is a completely independent array
	b := a

	// Print copied array
	fmt.Println("printing b:", b)

	// Print address of 'b'
	// This will be different from 'a' → proves it's a separate copy
	fmt.Printf("address of b: %p\n", &b)
}

func compare() {
	var a [2]int16 = [2]int16{2, 3}
	var b [2]int16 = [2]int16{2, 3}

	// Arrays can be compared directly using ==
	// because:
	// 1. Same type ([2]int16)
	// 2. Same length
	// 3. All elements are comparable

	if a == b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func iteration() {
	fmt.Println("for loop in arrays")

	var a [3]uint32 = [3]uint32{1, 2, 3}

	for i := 0; i < 3; i++ {
		fmt.Printf("index %d, value %d \n", i, a[i])
	}

	for i, v := range a {
		fmt.Println(i, v)
	}
}

func update(a [3]int) {
	// 'a' is a COPY of the original array
	// Any change here does NOT affect the caller

	a[0] = 0

	// prints the modified copy
	fmt.Println("inside update (copy):", a)
}

func update_pointer(a *[3]int) {
	// 'a' is a POINTER to the original array
	// Changes here WILL affect the caller

	a[0] = 0 // Go auto-dereferences (*a)[0] = 0
}

func modify() {
	// Create array (compiler infers size 3)
	arr := [...]int{1, 3, 5}

	// Passing array → full copy is made
	update(arr)

	// Original remains unchanged
	fmt.Println("after update (no pointer):", arr)

	// Passing pointer → no copy, same memory
	update_pointer(&arr)

	// Original is now modified
	fmt.Println("after update_pointer:", arr)
}

func memory_layout() {
	// Declare an array of 2 elements of type int8
	// int8 takes 1 byte → total array size = 2 bytes
	arr := [2]int8{1, 2}

	// Print address of the entire array
	// This points to the starting memory location of the array
	// Type: *[2]int8
	fmt.Printf("address of arr: %p \n", &arr)

	// Print address of first element
	// This is the same as the starting address of the array
	// Type: *int8
	fmt.Printf("address of arr[0]: %p \n", &arr[0])

	// Print address of second element
	// Since int8 = 1 byte, this address will be +1 from arr[0]
	fmt.Printf("address of arr[1]: %p \n", &arr[1])

	// Key Observations:
	// 1. &arr == &arr[0] → array starts at first element
	// 2. Elements are stored in contiguous memory
	// 3. Address difference depends on element size
	//    For int8 → +1 byte
	//    For int  → typically +8 bytes (on 64-bit system)

	// General formula:
	// address(arr[i]) = base_address + i * size_of_element
}
