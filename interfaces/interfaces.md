# Mastering Interfaces in Go

If you understand this file deeply and practice the examples, you'll have a **strong understanding of interfaces in Go**, which are one of the most important concepts in the language.

Interfaces are heavily used in:
- `net/http`
- `io`
- Kubernetes controllers
- database drivers
- middleware systems
- almost every real Go project

---

# Table of Contents

1. What Interfaces Are
2. Why Interfaces Exist
3. Declaring Interfaces
4. Implementing Interfaces
5. Implicit Implementation (Most Important Concept)
6. Interface Variables
7. Multiple Types Implementing Same Interface
8. Empty Interface (`interface{}`)
9. Type Assertions
10. Type Switches
11. Interfaces and Pointers
12. Method Sets
13. Real-world Example (`net/http`)
14. Interface Composition
15. Common Pitfalls
16. Hands-on Exercises
17. Quick Cheat Sheet

---

# What Interfaces Are

An interface defines **behavior**.

It tells:

```text
“If a type can do these methods, it satisfies this interface.”
```

Example:

```go
type Speaker interface {
	Speak()
}
```

This means:

```text
Any type having Speak() method becomes Speaker
```

---

# Why Interfaces Exist

Interfaces allow writing **generic behavior-based code**.

Without interfaces:

```text
code depends on concrete types ❌
```

With interfaces:

```text
code depends on behavior ✅
```

---

# Declaring Interfaces

```go
type Speaker interface {
	Speak()
}
```

Interface contains:
- method signatures only
- no implementation

---

# Implementing Interfaces

In Go, types implement interfaces automatically.

Example:

```go
type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof")
}
```

Now `Dog` satisfies:

```go
type Speaker interface {
	Speak()
}
```

because it has:
- `Speak()` method

---

# Implicit Implementation (VERY IMPORTANT)

Unlike Java/C++:

```text
No "implements" keyword exists in Go
```

Example:

```go
type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof")
}
```

`Dog` automatically satisfies `Speaker`.

---

# Full Example

```go
package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

func makeSpeak(s Speaker) {
	s.Speak()
}

func main() {
	d := Dog{}

	makeSpeak(d)
}
```

---

# Mental Model

```text
Interface = contract of behavior
```

NOT:

```text
Interface = inheritance
```

---

# Interface Variables

```go
var s Speaker
```

This variable can hold:

```text
Any value implementing Speaker
```

Example:

```go
s = Dog{}
```

---

# Multiple Types Implementing Same Interface

```go
type Speaker interface {
	Speak()
}

type Dog struct{}
type Cat struct{}

func (Dog) Speak() {
	fmt.Println("Woof")
}

func (Cat) Speak() {
	fmt.Println("Meow")
}
```

Both satisfy `Speaker`.

---

# Polymorphism in Go

```go
func makeSpeak(s Speaker) {
	s.Speak()
}
```

Can accept:
- Dog
- Cat
- Human
- anything implementing `Speak()`

---

# Empty Interface (`interface{}`)

```go
var x interface{}
```

Empty interface means:

```text
can hold ANY type
```

Example:

```go
x = 10
x = "hello"
x = true
```

Why?

Because every type satisfies:
- zero-method interface

---

# Type Assertions

Used to extract concrete value from interface.

```go
value := x.(int)
```

Example:

```go
var x interface{} = 10

v := x.(int)

fmt.Println(v)
```

---

# Safe Type Assertion

```go
v, ok := x.(int)
```

- `ok = true` → success
- `ok = false` → wrong type

---

# Type Switch

Used to check multiple concrete types.

```go
switch v := x.(type) {
case int:
	fmt.Println("int", v)

case string:
	fmt.Println("string", v)
}
```

---

# Interfaces and Pointers

Important concept.

---

## Value receiver

```go
type User struct{}

func (u User) Speak() {}
```

Both work:

```go
User{}
&User{}
```

---

## Pointer receiver

```go
func (u *User) Speak() {}
```

Only works with:

```go
&User{}
```

NOT:

```go
User{} // ❌
```

---

# Method Sets (IMPORTANT)

Method set determines interface satisfaction.

---

## Value receiver method

```go
func (u User) Speak()
```

Method belongs to:
- `User`
- `*User`

---

## Pointer receiver method

```go
func (u *User) Speak()
```

Method belongs only to:
- `*User`

---

# Real-world Example (`net/http`)

Core interface:

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

Any type implementing:

```go
ServeHTTP()
```

becomes HTTP handler.

---

# HandlerFunc Magic

Inside `net/http`:

```go
type HandlerFunc func(ResponseWriter, *Request)
```

And:

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

This is why normal functions can act as handlers.

---

# Interface Composition

Interfaces can embed other interfaces.

```go
type Reader interface {
	Read([]byte)
}

type Writer interface {
	Write([]byte)
}

type ReadWriter interface {
	Reader
	Writer
}
```

---

# Common Pitfalls

---

## 1. Nil interface confusion

```go
var s Speaker
fmt.Println(s == nil) // true
```

But:

```go
var d *Dog = nil
s = d

fmt.Println(s == nil) // false
```

Because interface stores:
- type
- value

---

## 2. Pointer receiver mismatch

```go
type User struct{}

func (u *User) Speak() {}
```

This fails:

```go
var s Speaker = User{} // ❌
```

Need:

```go
var s Speaker = &User{}
```

---

## 3. Overusing empty interface

```go
func process(x interface{})
```

Avoid unless necessary.

---

# Hands-on Exercises

1. Create `Animal` interface
2. Implement using Dog/Cat
3. Create empty interface variable
4. Practice type assertions
5. Practice type switches
6. Create custom HTTP handler
7. Experiment with pointer receivers

---

# Quick Cheat Sheet

```go
// interface
type Speaker interface {
	Speak()
}

// implementation
type Dog struct{}

func (d Dog) Speak() {}

// interface variable
var s Speaker = Dog{}

// empty interface
var x interface{}

// assertion
v := x.(int)

// safe assertion
v, ok := x.(int)

// type switch
switch v := x.(type) {}

// composition
type ReadWriter interface {
	Reader
	Writer
}
```


# Additional Section: Named Function Types and Interfaces

This section is EXTREMELY important because it explains one of the most beautiful concepts in Go:

```text
Functions are values
AND
methods can exist on function types
```

This is the foundation of:
- `net/http`
- middleware systems
- callbacks
- adapters
- functional patterns in Go

---

# What Is a Named Function Type?

A function can have its own named type.

Example:

```go
type MyFunc func()
```

This creates:

```text
a NEW named type whose underlying type is a function
```

---

# Basic Example

```go
package main

import "fmt"

type MyFunc func()

func hello() {
	fmt.Println("Hello")
}

func main() {
	var f MyFunc = hello

	f()
}
```

---

# Mental Model

```text
hello        → regular function
MyFunc       → named function type
f            → variable holding function value
```

---

# Functions Are First-class Values

In Go:

- functions can be assigned to variables
- functions can be passed to other functions
- functions can be returned
- functions can have named types

Example:

```go
func add(a int, b int) int {
	return a + b
}

var op func(int, int) int

op = add
```

---

# Methods on Function Types (VERY IMPORTANT)

This is where things become powerful.

Example:

```go
package main

import "fmt"

type MyFunc func()

func (f MyFunc) Call() {
	f()
}

func hello() {
	fmt.Println("Hello")
}

func main() {
	var f MyFunc = hello

	f.Call()
}
```

---

# Understanding What Happened

---

## Step 1

```go
type MyFunc func()
```

Created a named function type.

---

## Step 2

```go
func (f MyFunc) Call()
```

Attached method to function type.

---

## Step 3

```go
f()
```

Inside method:
- `f` itself is the function
- calling `f()` executes stored function

---

# Massive Insight

This means:

```text
functions can behave like objects
```

without classes.

---

# Relation to Interfaces

Function types can satisfy interfaces.

---

# Example

```go
package main

import "fmt"

type Caller interface {
	Call()
}

type MyFunc func()

func (f MyFunc) Call() {
	f()
}

func hello() {
	fmt.Println("Hello")
}

func main() {
	var c Caller

	c = MyFunc(hello)

	c.Call()
}
```

---

# Why This Works

Because:

```go
func (f MyFunc) Call()
```

means:

```text
MyFunc implements Caller
```

---

# Real-world Example: `net/http`

This is EXACTLY how Go's HTTP package works internally.

---

## Inside `net/http`

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

---

## Function Type

```go
type HandlerFunc func(ResponseWriter, *Request)
```

---

## Method on Function Type

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

---

# What This Means

A normal function:

```go
func hello(w http.ResponseWriter, r *http.Request)
```

can become:

```go
HandlerFunc(hello)
```

which now has:

```text
ServeHTTP()
```

and therefore satisfies:

```go
Handler
```

---

# Flow Visualization

```text
normal function
      ↓
converted into function type
      ↓
function type has methods
      ↓
satisfies interface
      ↓
used polymorphically
```

---

# Why This Pattern Is Powerful

It allows:
- simple APIs
- adapters
- middleware
- flexible architecture

without inheritance.

---

# Function Type vs Regular Function

| Feature | Regular Function | Named Function Type |
|---|---|---|
| Callable | ✅ | ✅ |
| Can have methods | ❌ | ✅ |
| Can satisfy interface | ❌ directly | ✅ |
| Has type identity | limited | strong |

---

# Another Example

```go
type Operation func(int, int) int

func (o Operation) Execute(a, b int) int {
	return o(a, b)
}

func add(a, b int) int {
	return a + b
}

func main() {
	op := Operation(add)

	fmt.Println(op.Execute(10, 20))
}
```

---

# Important Rule

Methods can only exist on:

```text
named types
```

Therefore:

```go
func (f func()) Test() {} // ❌ invalid
```

But:

```go
type MyFunc func()

func (f MyFunc) Test() {} // ✅
```

works.

---

# Realization

Go achieves many advanced patterns using only:

- functions
- named types
- methods
- interfaces

---

# Common Pitfall

People think:

```text
only structs can have methods ❌
```

Reality:

```text
ANY named type can have methods
```

including:
- ints
- strings
- slices
- maps
- functions

as long as they are named types.

---

# Hands-on Exercises

1. Create `MathFunc` type
2. Add method `Execute()`
3. Create interface `Runner`
4. Make function type satisfy interface
5. Recreate mini version of `http.HandlerFunc`

---

# Quick Cheat Sheet

```go
// named function type
type MyFunc func()

// method on function type
func (f MyFunc) Call() {
	f()
}

// interface
type Caller interface {
	Call()
}

// implementation
var c Caller = MyFunc(myFunction)
```

---

# Final Mental Model

```text
Named function types allow functions
to gain methods and satisfy interfaces.
```

This is one of the most elegant ideas in Go.

---

# One-line Takeaway

👉 “In Go, named function types can have methods, allowing functions themselves to satisfy interfaces.”
---

# Final Mental Model

```text
Interfaces describe behavior, not data.
```

Go programming is heavily:

```text
interface-driven design
```

---