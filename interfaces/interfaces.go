package main

import "fmt"

func main() {
	fmt.Println("---interfaces implementaton excercises----")
	example1()
	example2()
	type_assertion_switch()
	pointer_receiver()
	example3()
	example4()
	example5()
}

type Walker interface {
	Walk()
}

type Man struct {
	Name string
	Age  int
}

type Cat struct{}

func (m Man) Walk() {
	fmt.Println(m.Name, "is Walking")
}

func (c Cat) Walk() {
	fmt.Println("Cat is walking")

}
func Morning(w Walker) {
	w.Walk()
}

func example1() {
	m := Man{"sid", 23}
	c := Cat{}
	Morning(m)
	Morning(c)

}

//2nd example:

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Breadth * r.Length
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func example2() {
	r := Rectangle{1, 2}
	c := Circle{2}

	ss := []Shape{r, c}

	for _, val := range ss {
		fmt.Println(val.Area())
	}
}

//type asserstion and type switch

func check(t interface{}) {

	switch v := t.(type) {
	case int:
		fmt.Println("t is int and value is:", v)
	case string:
		fmt.Println("t is string and value: ", v)
	default:
		fmt.Println("default value:", v)
	}
}

func type_assertion_switch() {
	check(1)
	check("abc")

	type dummy interface{}

	i := dummy(1)

	val, ok := i.(int)

	fmt.Println(val, ok)
}

//value reciver and pointer receiver

type Writer interface {
	Write(string)
	Read(string)
}

type typist struct{}

func (t *typist) Write(data string) {
	fmt.Println("Printing data:", data)
}

func (t typist) Read(data string) {
	fmt.Println("data is:", data)
}

func pointer_receiver() {
	var w Writer

	w = &typist{}

	w.Write("DATA")

	// var r Writer = typist{} this doesn't work
	//r.Read("nothing")
}

//function type implements interface

type Sleeper interface {
	Sleep()
}

type Bus func()

func (b Bus) Sleep() {
	b()
}

func apsrtc() {
	fmt.Println("I am free for Women, only pallevelugu and express not luxury.")
}

func example3() {
	var bb Sleeper = Bus(apsrtc)

	bb.Sleep()
}

// composition example

type Jumpper interface {
	Jump()
}

type Swimmer interface {
	Swim()
}

type JumpSwimmer interface {
	Jumpper
	Swimmer
}

type Sid struct {
	Name string
}

func (s Sid) Jump() {
	fmt.Println(s.Name, "can Jump")
}

func (s Sid) Swim() {
	fmt.Println(s.Name, "can Swim")
}

func example4() {
	var n JumpSwimmer = Sid{"Siddhartha"}

	n.Jump()
	n.Swim()
}

// nill trap
type Errorer interface {
	Error()
}

type Uid struct{}

func (u *Uid) Error() {
	fmt.Println("UID error")
}

func generateError() Errorer {

	var u *Uid = nil

	return u
}

func example5() {

	err := generateError()

	fmt.Println("err == nil:", err == nil)

	fmt.Printf("Type: %T\n", err)

	err.Error()
}
