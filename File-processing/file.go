package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("---Heloo, File Processing module---")
	open_file()
	read_file()
	Open_and_Read()
	write_file()
	appending_text()
}

func open_file() {
	file, err := os.Open("text.txt")

	if err != nil {
		fmt.Println("Some ERROR occured", err)
	}
	defer file.Close()

	fmt.Println("file:", file)
}

func read_file() {
	file, _ := os.ReadFile("text.txt")

	fmt.Println(file)

}

func Open_and_Read() {
	file, _ := os.Open("text.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func write_file() {
	os.WriteFile("output.txt", []byte{72, 105}, 0644)
}

func appending_text() {
	file, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)

	defer file.Close()

	file.WriteString(" How are YOU? ")
}
