package main

import (
	"fmt"
)

func main() {
	fmt.Println("--------control statements in golang------------")
	if_else()
	switch_statement()
	for_loop()
	controlFlowAllInOne()
}

func if_else() {
	a := 7

	if a <= 5 {
		fmt.Println("give number a is less than 5")
	} else if a <= 10 {
		fmt.Println("given number a is less than 10 and greater than 5")
	} else {
		fmt.Println("given number a is greater than 10")
	}
}

func switch_statement() {
	b := 8

	switch b {
	case 1:
		fmt.Println("case 1 statement")
	case 2, 3, 4, 5:
		fmt.Println("second case statement")
	case 6, 7:
		fmt.Println("Third case statement")
	case 8:
		fmt.Println("fourth statement")
	}

	score := 320

	switch {
	case score > 300:
		fmt.Println("Score is more than 300")
	default:
		fmt.Println("score is less than 300")
	}
}

func for_loop() {
	I := 6

	for i := 0; i <= I; i++ {
		if i == I {
			fmt.Println("met the condition, breakin the loop")
			continue

		} else {
			fmt.Println(i)
			break
		}
	}
}

func controlFlowAllInOne() {
	fmt.Println("----- IF ELSE -----")
	a := 7

	if a <= 5 {
		fmt.Println("a <= 5")
	} else if a <= 10 {
		fmt.Println("5 < a <= 10")
	} else {
		fmt.Println("a > 10")
	}

	// ---------------------------------------

	fmt.Println("\n----- SWITCH (VALUE) -----")
	b := 3

	switch b {
	case 1:
		fmt.Println("one")
	case 2, 3:
		fmt.Println("two or three")
	default:
		fmt.Println("other")
	}

	// ---------------------------------------

	fmt.Println("\n----- SWITCH (CONDITION) -----")
	score := 85

	switch {
	case score >= 90:
		fmt.Println("Grade A")
	case score >= 75:
		fmt.Println("Grade B")
	default:
		fmt.Println("Grade C")
	}

	// ---------------------------------------

	fmt.Println("\n----- FOR LOOP (BASIC) -----")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// ---------------------------------------

	fmt.Println("\n----- FOR LOOP (WHILE STYLE) -----")
	i := 0
	for i < 3 {
		fmt.Println("while-style:", i)
		i++
	}

	// ---------------------------------------

	fmt.Println("\n----- INFINITE LOOP -----")
	count := 0
	for {
		fmt.Println("infinite loop iteration:", count)
		count++
		if count == 2 {
			break
		}
	}

	// ---------------------------------------

	fmt.Println("\n----- BREAK & CONTINUE -----")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // skip 2
		}
		if i == 4 {
			break // stop at 4
		}
		fmt.Println("value:", i)
	}

	// ---------------------------------------

	fmt.Println("\n----- NESTED LOOP -----")
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				break // breaks inner loop only
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	// ---------------------------------------

	fmt.Println("\n----- LABELED BREAK -----")
Outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				break Outer // breaks both loops
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

	// ---------------------------------------

	fmt.Println("\n----- RANGE LOOP (SLICE) -----")
	arr := []int{10, 20, 30}

	for i, v := range arr {
		fmt.Println("index:", i, "value:", v)
	}

	// ---------------------------------------

	fmt.Println("\n----- RANGE LOOP (MAP) -----")
	m := map[string]int{"a": 1, "b": 2}

	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}
}
