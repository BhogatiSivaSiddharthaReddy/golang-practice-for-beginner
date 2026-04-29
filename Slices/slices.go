package main

import "fmt"

func main() {
	slicePractice()
	sliceWTF()
	sliceParamDemo()
}

func slicePractice() {
	// 🔹 1. Nil slice
	// No memory allocated
	// ptr = nil, len = 0, cap = 0
	var s []int

	// 🔹 2. Empty slice
	// Memory exists (empty array), but len = 0
	// ptr ≠ nil, len = 0, cap = 0
	t := []int{}

	// 🔹 3. make with length
	// Allocates array of size 5
	// ptr → array, len = 5, cap = 5
	u := make([]int, 5)

	// 🔹 4. make with capacity
	// Allocates array of size 5
	// ptr → array, len = 0, cap = 5
	v := make([]int, 0, 5)

	fmt.Println("----- Initial State -----")

	// Printing:
	// len → number of elements
	// cap → total capacity
	// value → slice content
	// nil check → whether pointer is nil

	fmt.Printf("s -> len:%d cap:%d value:%v isNil:%t\n", len(s), cap(s), s, s == nil)
	fmt.Printf("t -> len:%d cap:%d value:%v isNil:%t\n", len(t), cap(t), t, t == nil)
	fmt.Printf("u -> len:%d cap:%d value:%v isNil:%t\n", len(u), cap(u), u, u == nil)
	fmt.Printf("v -> len:%d cap:%d value:%v isNil:%t\n", len(v), cap(v), v, v == nil)

	// Append behavior
	fmt.Println("\n----- Append Behavior -----")

	// Nil slice → allocates new array
	s = append(s, 1)

	// Empty slice → allocates new array
	t = append(t, 1)

	//
	u = append(u, 1)

	// Preallocated slice → uses existing capacity
	v = append(v, 1)

	fmt.Printf("s after append: %v (len:%d cap:%d)\n", s, len(s), cap(s))
	fmt.Printf("t after append: %v (len:%d cap:%d)\n", t, len(t), cap(t))
	fmt.Printf("v after append: %v (len:%d cap:%d)\n", v, len(v), cap(v))
	fmt.Printf("u after append: %v (len:%d cap:%d)\n", u, len(u), cap(u))

	// "" Shared memory example
	fmt.Println("\n----- Shared Memory Example -----")

	a := []int{10, 20, 30}
	b := a // copy slice header, NOT data

	b[0] = 999

	// Both change because they share same underlying array
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// "" Breaking sharing using copy
	fmt.Println("\n----- Copy to Avoid Sharing -----")

	c := make([]int, len(a))
	copy(c, a)

	c[0] = 111

	fmt.Println("a (unchanged):", a)
	fmt.Println("c (independent):", c)

	// "" Capacity growth demo
	fmt.Println("\n----- Capacity Growth -----")

	x := make([]int, 0)

	for i := 0; i < 10; i++ {
		x = append(x, i)
		fmt.Printf("len:%d cap:%d %v\n", len(x), cap(x), x)
	}
}

func sliceWTF() {
	// 🔹 Step 1: Create base array
	a := [3]int{1, 2, 3}

	// 🔹 Step 2: Create slice b from array
	// a[:1] → start=0, end=1
	// b = [1]
	// len = 1
	// cap = 3 (VERY IMPORTANT → extends till end of array)
	b := a[:1]

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("len(b):", len(b))
	fmt.Println("cap(b):", cap(b))

	fmt.Println("\n--- WTF Moment ---")

	// 🔹 Step 3: Extend slice beyond its length
	// This works because cap(b) = 3
	// So b[0:2] is VALID
	c := b[0:2]

	fmt.Println("c:", c)
	fmt.Println("len(c):", len(c))
	fmt.Println("cap(c):", cap(c))

	fmt.Println("\n--- Memory Sharing ---")

	// 🔹 Step 4: Modify c
	// This affects original array because all share same underlying array
	c[1] = 999

	fmt.Println("a after modifying c:", a)
	fmt.Println("b after modifying c:", b)
	fmt.Println("c after modifying c:", c)

	fmt.Println("\n--- Restricting Capacity (Fix) ---")

	// 🔹 Step 5: Restrict capacity using full slice expression
	// a[:1:1] → len=1, cap=1
	d := a[:1:1]

	fmt.Println("d:", d)
	fmt.Println("len(d):", len(d))
	fmt.Println("cap(d):", cap(d))

	fmt.Println("\n--- Now this will panic if uncommented ---")

	//  This will panic: slice bounds out of range
	// e := d[0:2]
	// fmt.Println(e)
}

// 🔹 Mutates existing elements (will reflect in caller)
func modifyElements(s []int) {
	s[0] = 999
}

// 🔹 Appends but DOES NOT return (may not reflect outside)
func appendNoReturn(s []int) {
	s = append(s, 100) // might reallocate → caller won't see it
	fmt.Println("inside appendNoReturn:", s)
}

// 🔹 Appends and RETURNS (correct way)
func appendWithReturn(s []int) []int {
	s = append(s, 200)
	return s
}

func sliceParamDemo() {
	// len=2, cap=2 → append will force reallocation
	s := make([]int, 2, 2)
	s[0], s[1] = 1, 2

	fmt.Println("original:", s)

	//  Mutation works (shared underlying array)
	modifyElements(s)
	fmt.Println("after modifyElements:", s)

	//  Append without return (lost if reallocated)
	appendNoReturn(s)
	fmt.Println("after appendNoReturn:", s)

	//  Correct append usage
	s = appendWithReturn(s)
	fmt.Println("after appendWithReturn:", s)
}
