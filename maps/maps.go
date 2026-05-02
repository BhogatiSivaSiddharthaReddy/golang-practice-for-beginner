package main

import "fmt"

func main() {
	fmt.Println("--------Learning about maps --------------")
	declaration()
	mapCRUD()
	iteration()
}

func declaration() {
	//  1. Nil map
	var mp map[string]string
	fmt.Println("nil map:", mp)

	//  2. Literal with values
	m1 := map[string]int{"num": 1}
	fmt.Println("literal map:", m1)

	//  3. Using make
	m2 := make(map[string]float64)
	fmt.Println("make map:", m2)

	//  4. Empty literal
	m3 := map[string]int{}
	fmt.Println("empty literal:", m3)

	//  5. Preallocated map
	m4 := make(map[string]int, 10)
	fmt.Println("preallocated map:", m4)

	//  6. Map with slice values
	m5 := map[string][]int{
		"nums": {1, 2, 3},
	}
	fmt.Println("map with slice:", m5)

	//  7. Map with struct values
	type User struct {
		Name string
	}

	m6 := map[string]User{
		"u1": {"Alice"},
	}
	fmt.Println("map with struct:", m6)
}

func mapCRUD() {
	// -------------------------------
	//  CREATE
	// -------------------------------
	// Initialize a map
	m := make(map[string]int)

	// Add key-value pairs
	m["apple"] = 10
	m["banana"] = 20

	fmt.Println("After CREATE:", m)

	// -------------------------------
	//  READ
	// -------------------------------
	// Direct read (returns zero value if key doesn't exist)
	val := m["apple"]
	fmt.Println("READ apple:", val)

	// Safe read using 'ok'
	val, ok := m["orange"]
	if ok {
		fmt.Println("orange exists:", val)
	} else {
		fmt.Println("orange not found")
	}

	// -------------------------------
	//  UPDATE
	// -------------------------------
	// Update existing key
	m["apple"] = 100
	fmt.Println("After UPDATE apple:", m)

	// -------------------------------
	//  DELETE
	// -------------------------------
	delete(m, "banana")
	fmt.Println("After DELETE banana:", m)

	// -------------------------------
	//  ITERATE
	// -------------------------------
	fmt.Println("Iterating map:")
	for k, v := range m {
		fmt.Printf("key=%s value=%d\n", k, v)
	}
}

func iteration() {
	//iteration of maps are unordered.
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}

	for k, v := range m {
		fmt.Println("key and value:", k, v)
	}

	for i, j := range m {
		fmt.Println("key and value:", i, j)
	}
}
