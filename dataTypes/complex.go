package main

import "fmt"

func complexNumbers() {
	fmt.Println("\n**** learning about COMPLEX data type ****")

	basicComplex()
	complexParts()
	complexOperations()
}

func basicComplex() {
	fmt.Println("\n**** Basic Complex ****")

	c1 := complex(2, 3) // 2 + 3i
	var c2 complex128 = 5 + 10i

	fmt.Println(c1)
	fmt.Println(c2)
}

func complexParts() {
	fmt.Println("\n**** Complex Parts ****")

	c := complex(4, 7)

	fmt.Println("Real:", real(c))
	fmt.Println("Imag:", imag(c))
}

func complexOperations() {
	fmt.Println("\n**** Complex Operations ****")

	c1 := complex(2, 3)
	c2 := complex(1, 1)

	fmt.Println("Add:", c1+c2)
	fmt.Println("Sub:", c1-c2)
	fmt.Println("Mul:", c1*c2)
	fmt.Println("Div:", c1/c2)
}
