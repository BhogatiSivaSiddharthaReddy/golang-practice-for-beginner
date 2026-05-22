package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("learning about go routines in golang")
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	go func() {
		fmt.Println("Hello from go routines from example 1")
	}()

	time.Sleep(time.Second)
}

func fib() func() int {

	fmt.Println("fib is called")

	a := 10

	return func() int {
		a = a + 10
		return a
	}

}

func example2() {
	fmt.Println("example two: Closures")
	f := fib
	f()
	f()

	g := fib()

	fmt.Println(g())
	fmt.Println(g())
}

func closures() func(a int) int {
	fmt.Println("Correct way")

	i := 10

	return func(a int) int {
		return i * a
	}
}

func example3() {
	fmt.Println("example three: correct way to share variable for closures")
	c := closures()
	fmt.Println(c(10))
}

func example4() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
