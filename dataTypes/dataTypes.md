# Mastering Basic Data Types in Go

If you understand this file deeply and practice the exercises, you'll have a **strong, practical, pro-level foundation** in Go basic data types.

---

## Table of Contents

1. [What Basic Data Types Are](#what-basic-data-types-are)
2. [Integer Types](#integer-types)
3. [Unsigned Integers](#unsigned-integers)
4. [Special Integer Types: byte & rune](#special-integer-types-byte--rune)
5. [Floating-Point Types](#floating-point-types)
6. [Complex Numbers](#complex-numbers)
7. [Boolean](#boolean)
8. [String](#string)
9. [Type Inference](#type-inference)
10. [Type Conversion](#type-conversion)
11. [Zero Values](#zero-values)
12. [Common Pitfalls](#common-pitfalls)
13. [Hands-on Exercises](#hands-on-exercises)
14. [Quick Cheat Sheet](#quick-cheat-sheet)

---

## What Basic Data Types Are

Basic data types are the **simplest building blocks** in Go.

They represent:

* Numbers (integers, floats, complex)
* Logical values (bool)
* Text (string)

Every variable you create in Go ultimately uses these.

👉 Important:
Basic types are **value types**, meaning values are copied during assignment.

---

## Integer Types

Integers store **whole numbers (no decimals)**.

### Why multiple integer types?

Because of **memory vs range trade-off**:

* Smaller type → less memory, smaller range
* Bigger type → more memory, larger range

### Types

```go
int8, int16, int32, int64, int
```

### Example

```go
var a int8 = 127
var b int16 = 32000
var c int32 = 2000000000
var d int64 = 9000000000000000000
var e int = 42
```

👉 Use `int` by default unless you need control over size.

---

## Unsigned Integers

Unsigned integers store **only positive values (0 and above)**.

### Types

```go
uint8, uint16, uint32, uint64, uint
```

### Example

```go
var u1 uint8 = 255
var u2 uint16 = 65000
var u3 uint32 = 4000000000
```

👉 Use when values can never be negative (like counts, sizes).

---

## Special Integer Types: byte & rune

### byte (alias of uint8)

Used for **raw data and ASCII**.

```go
var b byte = 'A'
fmt.Println(b) // 65
```

👉 Each character maps to a number internally.

---

### rune (alias of int32)

Used for **Unicode characters**.

```go
var r rune = '你'
fmt.Println(r)
fmt.Printf("%c", r)
```

👉 Key difference:

* byte → 1 byte (ASCII)
* rune → 4 bytes (Unicode)

---

## Floating-Point Types

Used for **decimal numbers**.

### Types

```go
float32, float64
```

### Example

```go
var f1 float32 = 3.14
var f2 float64 = 3.141592653589793
```

👉 `float64` is default because it provides better precision.

### Precision Issue

```go
var x float32 = 1.123456789
fmt.Println(x)
```

👉 Output is rounded due to binary storage.

---

## Complex Numbers

Used in mathematical computations.

```go
c := complex(2, 3)
fmt.Println(real(c))
fmt.Println(imag(c))
```

👉 Format: `a + bi`

---

## Boolean

Represents logical values.

```go
var isActive bool = true
isValid := false

fmt.Println(isActive && isValid)
```

👉 Used in conditions and control flow.

---

## String

Represents text.

### Properties

* Immutable
* UTF-8 encoded

### Example

```go
s := "Hello"
fmt.Println(len(s))
```

### Concatenation

```go
s1 := "Hello"
s2 := "World"
fmt.Println(s1 + " " + s2)
```

---

### Important: Byte vs Character Length

```go
s := "你好"
fmt.Println(len(s))         // bytes
fmt.Println(len([]rune(s))) // characters
```

👉 Common bug area in real-world systems.

---

## Type Inference

Go automatically determines types.

```go
a := 10
b := 3.14
c := true
d := "hello"
```

👉 Cleaner syntax, still strongly typed.

---

## Type Conversion

Go does **not allow implicit conversion**.

```go
var a int = 10
var b float64 = float64(a)

var c float64 = 9.7
var d int = int(c)
```

👉 Decimal part is truncated (not rounded).

---

## Zero Values

Variables always get a default value.

| Type   | Value |
| ------ | ----- |
| int    | 0     |
| float  | 0.0   |
| bool   | false |
| string | ""    |

### Example

```go
var a int
var b bool
fmt.Println(a, b)
```

---

## Common Pitfalls

### 1. Integer Overflow

```go
var x int8 = 127
x++
fmt.Println(x) // -128
```

### 2. Float Precision

```go
fmt.Println(0.1 + 0.2) // not exactly 0.3
```

### 3. String Length Confusion

```go
len("你好") != 2
```

---

## Hands-on Exercises

1. Declare all integer subtypes and print values
2. Demonstrate overflow in int8
3. Convert float to int and observe truncation
4. Compare byte vs rune in strings
5. Print ASCII values using byte
6. Handle Unicode string correctly
7. Show float precision issue

---

## Quick Cheat Sheet

```go
// integers
var a int = 10
var b int8 = 127

// unsigned
var u uint8 = 255

// float
f := 3.14

// complex
c := complex(1, 2)

// bool
flag := true

// string
s := "hello"

// conversion
float64(a)
int(f)

// rune & byte
r := '你'
b := byte('A')
```

---

Happy Learning 🚀