package main

import (
	"fmt"
	"math"
)

func floats() {
	fmt.Println("**** learning about FLOAT data type ****")
	basicFloat()
	floatPrecision()
	floatOperations()
	floatComparison()
	specialFloatValues()
}

func basicFloat() {
	fmt.Println("\n**** Basic Float ****")

	var f1 float32 = 3.14
	f2 := 3.1415926535 // default float64

	fmt.Println("float32:", f1)
	fmt.Println("float64:", f2)
}

func floatPrecision() {
	fmt.Println("\n**** Float Precision ****")

	a := 0.1
	b := 0.2

	fmt.Println("0.1 + 0.2 =", a+b)
}

func floatOperations() {
	fmt.Println("\n**** Float Operations ****")

	a := 10.5
	b := 2.5

	fmt.Println("Add:", a+b)
	fmt.Println("Sub:", a-b)
	fmt.Println("Mul:", a*b)
	fmt.Println("Div:", a/b)
}

func floatComparison() {
	fmt.Println("\n**** Float Comparison ****")

	a := 0.1 + 0.2
	b := 0.3

	fmt.Println("Direct compare:", a == b)

	// correct way
	epsilon := 1e-9
	fmt.Println("Approx equal:", math.Abs(a-b) < epsilon)
}

func specialFloatValues() {
	fmt.Println("\n**** Special Float Values ****")

	fmt.Println("Infinity:", math.Inf(1))
	fmt.Println("Negative Infinity:", math.Inf(-1))
	fmt.Println("NaN:", math.NaN())

	fmt.Println("NaN == NaN:", math.NaN() == math.NaN()) // false
}
