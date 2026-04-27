package main

import "fmt"

func main() {
	fmt.Println("This package will help you to learn basic data types")
	var info string = "Basic types are value types, meaning values are copied during assignment. It means when you define a basic type for a variable and assign a value to that variable later you assign this defined variable to another variable instead of referencing the memory addres it copies the data and assign to it."
	fmt.Println(info)
	integers()
	bytes()
	runes()
	stringsDemo()
}
