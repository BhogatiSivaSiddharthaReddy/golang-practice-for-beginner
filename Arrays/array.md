# Mastering Arrays in Go (Complete Guide)

If you understand this file deeply, you’ll have a **strong foundation in memory and data structures**, which is essential before learning slices.

---

## Table of Contents

1. What is an Array?
2. Why Arrays Exist in Go
3. Declaration & Initialization
4. Type System & Array Identity
5. Accessing & Updating Elements
6. Iteration Patterns
7. Value Semantics (Critical)
8. Arrays in Functions
9. Pointer to Arrays
10. Multi-Dimensional Arrays
11. Memory Layout & Internals
12. Arrays vs Slices
13. Performance Considerations
14. Comparability
15. Zero Values
16. Arrays of Complex Types
17. Copy vs Reference Deep Dive
18. Common Pitfalls
19. When to Use Arrays
20. Hands-on Exercises
21. Quick Cheat Sheet

---

## 1. What is an Array?

An array is a **fixed-size collection of elements of the same type stored in contiguous memory**.

```go
var arr [3]int
```

Key properties:

* Fixed size
* Same type
* Stored sequentially in memory

---

## 2. Why Arrays Exist in Go

Arrays provide:

* Predictable memory layout
* Compile-time size guarantees
* Foundation for slices

---

## 3. Declaration & Initialization

### Basic declaration

```go
var arr [3]int
```

### With values

```go
arr := [3]int{1, 2, 3}
```

### Partial initialization

```go
arr := [5]int{1, 2}
// Output: [1 2 0 0 0]
```

### Compiler decides size

```go
arr := [...]int{1, 2, 3}
```

---

## 4. Type System & Array Identity

Array size is part of the type:

```go
var a [3]int
var b [4]int

// a = b ❌ error
```

---

## 5. Accessing & Updating Elements

```go
arr := [3]int{10, 20, 30}

fmt.Println(arr[0])
arr[1] = 99
```

---

## 6. Iteration Patterns

### Classic loop

```go
for i := 0; i < len(arr); i++ {
	fmt.Println(arr[i])
}
```

### Range loop

```go
for i, v := range arr {
	fmt.Println(i, v)
}
```

---

## 7. Value Semantics (Critical)

Arrays are copied:

```go
a := [3]int{1,2,3}
b := a

b[0] = 100

fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

---

## 8. Arrays in Functions

```go
func modify(a [3]int) {
	a[0] = 100
}
```

Original array is unchanged.

---

## 9. Pointer to Arrays

```go
func modify(a *[3]int) {
	a[0] = 100
}
```

Now original array is modified.

---

## 10. Multi-Dimensional Arrays

```go
matrix := [2][3]int{
	{1,2,3},
	{4,5,6},
}
```

---

## 11. Memory Layout & Internals

Arrays use **contiguous memory**:

```
| 1 | 2 | 3 |
```

This makes access fast and predictable.

---

## 12. Arrays vs Slices

| Feature | Array     | Slice       |
| ------- | --------- | ----------- |
| Size    | Fixed     | Dynamic     |
| Copy    | Full copy | Shared data |
| Usage   | Rare      | Common      |

---

## 13. Performance Considerations

Large arrays are expensive to copy:

```go
func f(a [100000]int) {}
```

Better:

```go
func f(a *[100000]int) {}
```

---

## 14. Comparability

Arrays can be compared directly:

```go
a := [3]int{1,2,3}
b := [3]int{1,2,3}

fmt.Println(a == b) // true
```

---

## 15. Zero Values

```go
var arr [3]int
```

Default:

```
[0 0 0]
```

---

## 16. Arrays of Complex Types

```go
type Person struct {
	Name string
}

arr := [2]Person{
	{"Alice"},
	{"Bob"},
}
```

---

## 17. Copy vs Reference Deep Dive

```go
a := [3]int{1,2,3}
b := a

b[0] = 99
```

👉 `a` remains unchanged → deep copy

---

## 18. Common Pitfalls

### Different sizes = different types

```go
[3]int != [4]int
```

### Expensive copying

```go
func f(a [1000]int) {}
```

### Not dynamic

```go
arr := [3]int{1,2,3}
// cannot resize
```

---

## 19. When to Use Arrays

Use arrays when:

* size is fixed
* memory must be predictable
* performance-critical code

In most real-world Go code → use **slices**

---

## 20. Hands-on Exercises

### Beginner

1. Declare and print an array
2. Update values
3. Find sum

### Intermediate

4. Reverse array
5. Find max/min
6. Copy and verify independence

### Advanced

7. Pass array to function
8. Modify using pointer
9. Work with 2D arrays

---

## 21. Quick Cheat Sheet

```go
// declaration
var arr [3]int

// initialization
arr := [3]int{1,2,3}

// length
len(arr)

// iteration
for i, v := range arr {}

// copy
b := arr

// pointer
func f(a *[3]int) {}
```

---

## Final Tip

Arrays are foundational.

Master them to understand:

* memory
* value semantics
* performance
