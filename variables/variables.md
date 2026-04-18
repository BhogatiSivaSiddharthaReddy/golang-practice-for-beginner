# Mastering Variables in Go

If you understand this file deeply and practice the exercises, you'll have a **strong, practical, pro-level foundation** in Go variables.

## Table of Contents

1. [What a Variable Is](#what-a-variable-is)
2. [Ways to Declare Variables](#ways-to-declare-variables)
3. [Type Inference vs Explicit Types](#type-inference-vs-explicit-types)
4. [Zero Values (Very Important)](#zero-values-very-important)
5. [Reassignment and Type Safety](#reassignment-and-type-safety)
6. [Multiple Variables and Tuple Assignment](#multiple-variables-and-tuple-assignment)
7. [Scope Rules](#scope-rules)
8. [Shadowing (`:=`) Pitfalls](#shadowing--pitfalls)
9. [Package-Level Variables](#package-level-variables)
10. [Pointers and Variables](#pointers-and-variables)
11. [Reference-like Types and Variables](#reference-like-types-and-variables)
12. [Variables in Loops and Closures](#variables-in-loops-and-closures)
13. [Concurrency and Variable Safety](#concurrency-and-variable-safety)
14. [Constants vs Variables](#constants-vs-variables)
15. [Naming and Best Practices](#naming-and-best-practices)
16. [Common Interview-Style Patterns](#common-interview-style-patterns)
17. [Hands-on Exercises](#hands-on-exercises)
18. [Quick Cheat Sheet](#quick-cheat-sheet)

---

## What a Variable Is

A **variable** is a named storage location that holds a value.

In Go, every variable has:
- a **name** (`count`)
- a **type** (`int`, `string`, `[]int`, etc.)
- a **value** (for example `42`)

```go
var count int = 42
```

Think of it like:
- `count` = label
- `int` = kind of data allowed
- `42` = current stored data

---

## Ways to Declare Variables

### 1) `var name type = value` (explicit type + value)

```go
var age int = 25
var name string = "Asha"
```

Use when you want to be very explicit.

### 2) `var name = value` (type inferred)

```go
var age = 25       // inferred as int
var name = "Asha" // inferred as string
```

Compiler infers the type from value.

### 3) `var name type` (zero value assigned)

```go
var age int        // 0
var name string    // ""
var ok bool        // false
```

Great when you need declaration now, assignment later.

### 4) Short declaration `:=` (inside functions only)

```go
func main() {
    age := 25
    name := "Asha"
    _ = age
    _ = name
}
```

Rules:
- only inside functions
- at least one variable on the left side must be new

---

## Type Inference vs Explicit Types

```go
x := 10        // int
pi := 3.14     // float64
ok := true     // bool
msg := "hello" // string
```

Explicit type can be useful for control:

```go
var small int8 = 100
var price float32 = 19.99
```

If not explicit, numeric defaults are typically:
- integer literal -> `int`
- floating literal -> `float64`
- rune literal -> `rune` (`int32`)

---

## Zero Values (Very Important)

Go never leaves variables uninitialized.

| Type | Zero Value |
|------|------------|
| `int`, `int64`, `float64` | `0` / `0.0` |
| `bool` | `false` |
| `string` | `""` |
| pointer | `nil` |
| slice | `nil` |
| map | `nil` |
| channel | `nil` |
| function | `nil` |
| interface | `nil` |
| struct | all fields zero-valued |

Example:

```go
package main

import "fmt"

func main() {
    var n int
    var s string
    var b bool
    var nums []int

    fmt.Printf("n=%d, s=%q, b=%v, nums=%v, nums==nil:%v\n", n, s, b, nums, nums == nil)
}
```

---

## Reassignment and Type Safety

Go is statically typed.

```go
x := 10
x = 20      // okay
// x = "hi" // compile error: cannot use string as int
```

You may convert explicitly:

```go
var i int = 42
var f float64 = float64(i)
```

---

## Multiple Variables and Tuple Assignment

### Declare multiple variables

```go
var a, b, c int = 1, 2, 3
x, y := "go", 1.23
```

### Swap values (idiomatic)

```go
a, b := 10, 20
a, b = b, a
```

### Receive multiple returns

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    n, err := strconv.Atoi("123")
    fmt.Println(n, err) // 123 <nil>
}
```

### Ignore values with blank identifier `_`

```go
n, _ := strconv.Atoi("123") // ignore error (only when safe)
_ = n
```

---

## Scope Rules

### 1) Block scope

A variable exists only inside the nearest `{ ... }` block where declared.

```go
func demo() {
    x := 10
    {
        y := 20
        _ = x
        _ = y
    }
    // _ = y // compile error: undefined y
}
```

### 2) Function scope

Variables declared at top-level of function are visible throughout function after declaration.

### 3) Package scope

Variables declared outside functions are visible throughout package (subject to export rules).

---

## Shadowing (`:=`) Pitfalls

Shadowing means creating a **new variable** with same name in inner scope.

```go
package main

import "fmt"

func main() {
    x := 10
    if true {
        x := 99 // shadows outer x
        fmt.Println("inner x:", x)
    }
    fmt.Println("outer x:", x)
}
```

Output:

```text
inner x: 99
outer x: 10
```

### Common bug with `err`

```go
f, err := openFile()
if err != nil {
    return err
}

if data, err := readData(f); err != nil { // new err scoped to if
    return err
} else {
    _ = data
}

// here, original err is still the outer one
```

Be careful when using `:=` in nested scopes.

---

## Package-Level Variables

```go
package main

import "fmt"

var appName = "BillingService" // package-level
var maxConn int                 // zero value: 0

func main() {
    fmt.Println(appName, maxConn)
}
```

Notes:
- Package-level variables are initialized before `main()`.
- Initialization order matters across files in same package.

---

## Pointers and Variables

A pointer stores the address of another variable.

```go
package main

import "fmt"

func main() {
    x := 10
    p := &x        // p points to x
    *p = 99        // update x via pointer

    fmt.Println(x) // 99
}
```

Why this matters:
- pass large structs efficiently
- allow function to modify caller's data

```go
func setZero(n *int) {
    *n = 0
}
```

---

## Reference-like Types and Variables

These types behave differently from plain values:
- `slice`
- `map`
- `channel`
- `function`

### Slice example

```go
a := []int{1, 2, 3}
b := a
b[0] = 99
fmt.Println(a) // [99 2 3] (shares underlying array)
```

### Map example

```go
m1 := map[string]int{"go": 1}
m2 := m1
m2["go"] = 2
fmt.Println(m1["go"]) // 2
```

These variables copy descriptors/headers, not deep data.

---

## Variables in Loops and Closures

Classic closure bug: capturing loop variable incorrectly.

```go
package main

import "fmt"

func main() {
    fns := []func(){}

    for i := 0; i < 3; i++ {
        i := i // create new per-iteration variable (safe)
        fns = append(fns, func() { fmt.Println(i) })
    }

    for _, fn := range fns {
        fn()
    }
}
```

Without the `i := i` line, behavior can be surprising in many closure scenarios.

---

## Concurrency and Variable Safety

If multiple goroutines access the same variable, and at least one writes, you need synchronization.

### Unsafe (data race)

```go
// counter++ from multiple goroutines without protection => race
```

### Safe with mutex

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    counter := 0
    wg := sync.WaitGroup{}

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }

    wg.Wait()
    fmt.Println(counter)
}
```

Use `go run -race` to detect race conditions.

---

## Constants vs Variables

Use `const` for values that should never change.

```go
const Pi = 3.14159
var radius = 10.0
area := Pi * radius * radius
```

Differences:
- `const` cannot be reassigned
- constants can be untyped and more flexible in expressions

```go
const n = 10 // untyped constant
var i int64 = n // okay
```

---

## Naming and Best Practices

- Use short names for short scopes: `i`, `n`, `err`
- Use meaningful names for wider scope: `retryLimit`, `httpTimeout`
- Avoid overly generic names: `data`, `temp`, `value1`
- Keep variables as close as possible to first use
- Prefer `:=` in local code, but avoid accidental shadowing
- Minimize package-level mutable state
- Use `const` whenever possible

---

## Common Interview-Style Patterns

### Pattern 1: Two-value map lookup

```go
m := map[string]int{"go": 1}
v, ok := m["go"]
fmt.Println(v, ok) // 1 true
```

### Pattern 2: Type assertion with `ok`

```go
var any interface{} = "hello"
s, ok := any.(string)
fmt.Println(s, ok)
```

### Pattern 3: Comma-ok with channel receive

```go
ch := make(chan int)
close(ch)
v, ok := <-ch
fmt.Println(v, ok) // 0 false
```

---

## Hands-on Exercises

Practice these to become truly confident:

1. Declare one variable each using all 4 declaration styles.
2. Print zero values for `int`, `bool`, `string`, `[]int`, `map[string]int`.
3. Write a function that swaps two integers using tuple assignment.
4. Create a shadowing bug intentionally; then fix it.
5. Write a function that updates an integer via pointer.
6. Create 100 goroutines incrementing a shared counter; first unsafe, then safe with mutex.
7. Build a small config struct and compare passing by value vs pointer.

---

## Quick Cheat Sheet

```go
// declarations
var a int = 10
var b = 20
var c int

// short declaration (inside funcs only)
d := 30

// multiple
a1, b1 := 1, 2

// swap
a1, b1 = b1, a1

// blank identifier
x, _ := someFunc()

// pointer
p := &a
*p = 99

// map lookup
v, ok := myMap["k"]

// type assertion
s, ok := any.(string)
```

---

## Final Notes to Reach Pro Level

To become truly pro with Go variables:
- write code daily (small focused snippets)
- read compile errors carefully (they teach types/scopes fast)
- run `go vet` and `go run -race`
- review your code for shadowing and unnecessary global state

If you'd like, next we can create a **practice workbook** with 25 variable-focused problems and solutions.
