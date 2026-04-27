package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func stringsDemo() {
	fmt.Println("**** learning about STRING data type ****")

	basicString()
	stringLength()
	stringIndexing()
	stringIter()
	stringModification()
	stringConcatenation()
	stringComparison()
	stringConversion()
	stringFunctions()
	stringBugs()
}

func basicString() {
	fmt.Println("\n**** Basic String ****")

	s := "Hello"
	var s2 string = "World"

	fmt.Println(s, s2)
}

func stringLength() {
	fmt.Println("\n**** String Length ****")

	s := "A你"

	fmt.Println("String:", s)
	fmt.Println("Byte length:", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
}

func stringIndexing() {
	fmt.Println("\n**** String Indexing ****")

	s := "A你"

	for i := 0; i < len(s); i++ {
		fmt.Printf("byte[%d] = %d\n", i, s[i])
	}
}

func stringIter() {
	fmt.Println("\n**** String Iteration (rune safe) ****")

	s := "Go你"

	for i, r := range s {
		fmt.Printf("index=%d char=%c value=%d\n", i, r, r)
	}
}

func stringModification() {
	fmt.Println("\n**** String Modification ****")

	s := "hello ₹"

	runes := []rune(s)
	fmt.Println(runes)
	runes[0] = 'H'

	s = string(runes)

	fmt.Println("Modified:", s)
}

func stringConcatenation() {
	fmt.Println("\n**** String Concatenation ****")

	s1 := "Hello"
	s2 := "World"

	fmt.Println(s1 + " " + s2)

	// efficient way
	var builder strings.Builder
	builder.WriteString(s1)
	builder.WriteString(" ")
	builder.WriteString(s2)

	fmt.Println("Builder:", builder.String())
}

func stringComparison() {
	fmt.Println("\n**** String Comparison ****")

	fmt.Println("abc == abc:", "abc" == "abc")
	fmt.Println("abc < abd :", "abc" < "abd")
}

func stringConversion() {
	fmt.Println("\n**** String Conversion ****")

	s := "hello"

	bytes := []byte(s)
	fmt.Println("To bytes:", bytes)

	newStr := string(bytes)
	fmt.Println("Back to string:", newStr)
}

func stringFunctions() {
	fmt.Println("\n**** String Functions ****")

	fmt.Println(strings.Contains("hello", "he"))
	fmt.Println(strings.ToUpper("go"))
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.Split("a,b,c", ","))
	fmt.Println(strings.Join([]string{"a", "b"}, ","))
}

func stringBugs() {
	fmt.Println("\n**** Common String Bugs ****")

	s := "你好"

	fmt.Println("len (wrong for chars):", len(s))
	fmt.Println("correct rune count:", utf8.RuneCountInString(s))

	fmt.Println("\nByte iteration (wrong for chars):")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
