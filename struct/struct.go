package main

import (
	"fmt"
)

func main() {
	fmt.Println("------Learning about structs---------")
	declaring()
}

func declaring() {
	fmt.Println("Different ways to define structs")

	// 1. Inline struct (declare then assign)
	var u struct {
		Name string
		Age  int
	}

	u = struct {
		Name string
		Age  int
	}{
		"Sid",
		12,
	}
	fmt.Println(u, "---")

	// 2. Anonymous struct
	user := struct {
		Name string
		Age  int
	}{
		"Sid",
		24,
	}
	fmt.Println(user, "---")

	// 3. Named struct
	type person struct {
		name string
		age  int
	}

	// 4. Named struct initialization
	var p person
	p = person{"sid", 25}
	fmt.Println(p, "----")

	// 5. Direct initialization (idiomatic)
	p2 := person{"ram", 30}
	fmt.Println(p2, "----")

	// 6. Named field initialization (best)
	p3 := person{
		name: "asha",
		age:  22,
	}
	fmt.Println(p3, "----")

	// 7. Pointer to struct
	ptr := &person{"john", 40}
	fmt.Println(ptr, "----")

	// 8. Zero value struct
	var empty person
	fmt.Println(empty, "----")
}
