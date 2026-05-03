
# Mastering Structs in Go

If you understand this file deeply and practice the examples, you'll have a **strong, practical understanding of structs in Go**, which are the foundation for modeling real-world data.

---

## Table of Contents

1. What a Struct Is
2. Declaring Structs
3. Creating Struct Instances
4. Accessing and Updating Fields
5. Zero Values in Structs
6. Structs and Pointers
7. Structs as Function Parameters
8. Anonymous Structs
9. Nested Structs
10. Struct Tags (Very Important)
11. Methods on Structs
12. Value vs Pointer Receivers
13. Struct Comparison Rules
14. Embedded Structs (Composition)
15. JSON and Structs (Real-world Use)
16. Common Pitfalls
17. Hands-on Exercises
18. Quick Cheat Sheet

---

## What a Struct Is

A **struct** is a collection of fields grouped together.

It helps represent real-world entities.

```go
type User struct {
	Name string
	Age  int
}
````

Think of it like:

```
User = blueprint
fields = properties
```

---

## Declaring Structs

```go
type User struct {
	Name string
	Age  int
}
```

---

## Creating Struct Instances

### 1. Using struct literal

```go
u := User{
	Name: "Sid",
	Age:  25,
}
```

---

### 2. Without field names (order matters)

```go
u := User{"Sid", 25}
```

---

### 3. Using `new`

```go
u := new(User) // returns *User
```

---

## Accessing and Updating Fields

```go
fmt.Println(u.Name)

u.Age = 30
```

---

## Zero Values in Structs

If not initialized:

```go
var u User
```

Then:

| Field  | Value |
| ------ | ----- |
| string | ""    |
| int    | 0     |
| bool   | false |

---

## Structs and Pointers

```go
u := User{Name: "Sid", Age: 25}

p := &u

p.Age = 30
```

👉 Go automatically dereferences pointer

---

## Structs as Function Parameters

### By value (copy)

```go
func update(u User) {
	u.Age = 50
}
```

👉 original not changed

---

### By pointer

```go
func update(u *User) {
	u.Age = 50
}
```

👉 original updated

---

## Anonymous Structs

```go
u := struct {
	Name string
	Age  int
}{
	Name: "Sid",
	Age:  25,
}
```

👉 Useful for temporary data

---

## Nested Structs

```go
type Address struct {
	City string
}

type User struct {
	Name    string
	Address Address
}
```

---

## Struct Tags (Very Important)

Used for encoding/decoding

```go
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

---

## Methods on Structs

```go
func (u User) greet() {
	fmt.Println("Hello", u.Name)
}
```

---

## Value vs Pointer Receivers

### Value receiver

```go
func (u User) update() {
	u.Age = 50
}
```

👉 copy, no change

---

### Pointer receiver

```go
func (u *User) update() {
	u.Age = 50
}
```

👉 modifies original

---

## Struct Comparison Rules

Structs can be compared if all fields are comparable:

```go
a := User{"Sid", 25}
b := User{"Sid", 25}

fmt.Println(a == b) // true
```

---

## Embedded Structs (Composition)

```go
type Address struct {
	City string
}

type User struct {
	Name string
	Address
}
```

Access:

```go
u.City
```

---

## JSON and Structs (Real-world Use)

```go
data := `{"name":"Sid","age":25}`

var u User
json.Unmarshal([]byte(data), &u)
```

---

## Common Pitfalls

---

### ❌ Forgetting pointer for updates

```go
func update(u User) {
	u.Age = 50 // no effect
}
```

---

### ❌ Using unexported fields with JSON

```go
type User struct {
	name string // won't work with JSON
}
```

---

### ❌ Comparing structs with slices/maps

```go
type A struct {
	Data []int
}

// fmt.Println(a == b) ❌ not allowed
```

---

## Hands-on Exercises

1. Create a `Book` struct and print values
2. Write function to update struct using pointer
3. Create nested struct (User + Address)
4. Add method to struct
5. Parse JSON into struct

---

## Quick Cheat Sheet

```go
// define
type User struct {
	Name string
	Age  int
}

// create
u := User{"Sid", 25}

// pointer
p := &u
p.Age = 30

// method
func (u User) greet() {}

// pointer method
func (u *User) update() {}

// anonymous
u := struct{ Name string }{"Sid"}

// nested
type A struct {
	B struct{}
}
```

---

## Final Mental Model

```
Struct = custom data type grouping related fields
```

