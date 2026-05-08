# Mastering Methods in Go

Methods are one of the most important features in Go because they allow behavior to be attached to types.

If you understand this file deeply, you'll understand:
- methods
- receivers
- pointer receivers
- method sets
- interfaces
- Go's object-oriented style

Methods are heavily used in:
- structs
- interfaces
- net/http
- Kubernetes controllers
- database clients
- middleware systems

---

# Table of Contents

1. What Methods Are
2. Methods vs Functions
3. Receiver Syntax
4. Value Receivers
5. Pointer Receivers
6. Automatic Dereferencing
7. Automatic Addressing
8. Methods on Named Types
9. Methods on Function Types
10. Method Sets
11. Methods and Interfaces
12. Methods vs Struct Functions
13. Common Pitfalls
14. Hands-on Exercises
15. Quick Cheat Sheet

---

# What Methods Are

A method is simply:

```text
a function attached to a type
```

---

# Basic Syntax

```go
func (receiver Type) methodName() {
	// logic
}
```

Example:

```go
type User struct {
	Name string
}

func (u User) greet() {
	fmt.Println("Hello", u.Name)
}
```

---

# Calling Methods

```go
u := User{"Sid"}

u.greet()
```

---

# Methods vs Functions

---

## Normal Function

```go
func greet(u User) {
	fmt.Println(u.Name)
}
```

Called like:

```go
greet(u)
```

---

## Method

```go
func (u User) greet() {
	fmt.Println(u.Name)
}
```

Called like:

```go
u.greet()
```

---

# Mental Model

Method:

```go
u.greet()
```

is roughly equivalent to:

```go
greet(u)
```

---

# Receiver Syntax

Receiver appears between:

```go
func (...) method()
```

Example:

```go
func (u User) greet()
```

Here:
- `u` → receiver variable
- `User` → receiver type

---

# Value Receivers

---

## Example

```go
type User struct {
	Name string
	Age  int
}

func (u User) update() {
	u.Age = 50
}
```

---

# Important

Value receiver gets COPY of struct.

Original does NOT change.

---

# Example

```go
u := User{"Sid", 25}

u.update()

fmt.Println(u.Age)
```

Output:

```text
25
```

---

# Pointer Receivers

---

## Example

```go
func (u *User) update() {
	u.Age = 50
}
```

---

# Important

Pointer receiver gets pointer to original struct.

Original changes.

---

# Example

```go
u := User{"Sid", 25}

u.update()

fmt.Println(u.Age)
```

Output:

```text
50
```

---

# When to Use Pointer Receivers

Use pointer receivers when:
- modifying struct
- avoiding copies
- struct is large
- consistency across methods

---

# Automatic Dereferencing

Go automatically dereferences pointers for field access and method calls.

---

# Example

```go
u := User{"Sid", 25}

p := &u

fmt.Println(p.Name)
```

Internally:

```go
(*p).Name
```

---

# Automatic Addressing

Go also automatically takes address for method calls.

---

# Example

```go
func (u *User) greet() {}

u := User{}

u.greet()
```

Internally:

```go
(&u).greet()
```

---

# Important

This automatic conversion works only:
- for method calls
- for struct field access

NOT for interface satisfaction.

---

# Methods on Named Types

Methods can exist on ANY named type.

Not only structs.

---

# Example: Named Int Type

```go
type MyInt int

func (m MyInt) double() MyInt {
	return m * 2
}
```

---

# Example: Named Slice Type

```go
type Numbers []int

func (n Numbers) sum() int {
	total := 0

	for _, v := range n {
		total += v
	}

	return total
}
```

---

# Methods on Function Types

Very important concept.

---

# Example

```go
type MyFunc func()

func (f MyFunc) Call() {
	f()
}
```

---

# Why Powerful?

This allows:
- functions to satisfy interfaces
- adapters
- middleware patterns

Used heavily in:
- net/http

---

# Real-world Example (`net/http`)

---

# Interface

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

---

# Function Type

```go
type HandlerFunc func(ResponseWriter, *Request)
```

---

# Method on Function Type

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

---

# Massive Insight

Normal functions become interface implementations.

---

# Method Sets

Method set determines:
- interface satisfaction
- accessible methods

---

# Value Receiver Methods

```go
func (u User) walk()
```

Belong to:
- `User`
- `*User`

---

# Pointer Receiver Methods

```go
func (u *User) speak()
```

Belong ONLY to:
- `*User`

---

# Methods and Interfaces

---

# Interface

```go
type Speaker interface {
	speak()
}
```

---

# Pointer Receiver Example

```go
type User struct{}

func (u *User) speak() {}
```

---

# This FAILS

```go
var s Speaker = User{}
```

because:
- `speak()` belongs only to `*User`

---

# This WORKS

```go
var s Speaker = &User{}
```

---

# Value Receiver Example

```go
func (u User) speak() {}
```

Now BOTH work:

```go
var s1 Speaker = User{}
var s2 Speaker = &User{}
```

---

# Important Rule

```text
For interface satisfaction:
- value receiver methods belong to T and *T
- pointer receiver methods belong only to *T
```

---

# Methods vs Struct Functions

---

## Method

```go
func (u User) greet()
```

Behavior attached to type.

---

## Function

```go
func greet(u User)
```

Standalone logic.

---

# Common Pitfalls

---

## 1. Confusing automatic method calls with interface satisfaction

This works:

```go
u.speak()
```

because Go auto-addresses.

But this fails:

```go
var s Speaker = u
```

if method has pointer receiver.

---

## 2. Forgetting value receiver copies data

```go
func (u User) update()
```

modifies copy only.

---

## 3. Using pointer receivers inconsistently

Avoid mixing:
- some methods value receiver
- some pointer receiver

unless necessary.

---

# Hands-on Exercises

1. Create User struct
2. Add greet() method
3. Add updateAge() pointer receiver
4. Create named int type with method
5. Create named slice type with method
6. Create function type with method
7. Create interface implemented by method

---

# Quick Cheat Sheet

```go
// value receiver
func (u User) greet()

// pointer receiver
func (u *User) update()

// method call
u.greet()

// automatic addressing
u.update()

// named type method
type MyInt int

func (m MyInt) double()

// function type method
type MyFunc func()

func (f MyFunc) Call()
```

---

# Final Mental Model

```text
Methods attach behavior to types.
```

Go achieves object-oriented behavior using:
- methods
- interfaces
- composition

without classes.

---

# Deep Insight

```text
Receiver type determines:
- mutation behavior
- method sets
- interface satisfaction
```

---

# One-line Takeaway

👉 “Methods in Go are functions attached to types, enabling behavior-oriented programming.”