package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Heloo, Hola standarad otuput")

	fmt.Fprintln(os.Stderr, ("printing to error otuput"))
}
