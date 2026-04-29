# Maps in Go — From Basics to Advanced (Friendly Guide)

This guide starts simple and gradually builds to advanced concepts. Read top → bottom.

---

## Table of Contents

1. What is a Map?
2. Creating Maps (3 ways)
3. Read, Write, Delete
4. Check if Key Exists (IMPORTANT)
5. Iteration (order warning)
6. Nil vs Empty Map
7. Maps are Reference Types
8. Map Internals (gentle intro)
9. Performance Tips (simple)
10. Advanced Topics (overview)
11. Concurrency (must know)
12. Common Pitfalls
13. Practice Snippets
14. Cheat Sheet

---

## 1. What is a Map?

A map stores **key → value** pairs.

```go
m := map[string]int{
	"apple": 10,
	"banana": 20,
}
```

---

## 2. Creating Maps

### A. Using make (most common)

```go
m := make(map[string]int)
```

### B. Using literal

```go
m := map[string]int{"a": 1, "b": 2}
```

### C. Nil map (read-only safe)

```go
var m map[string]int
```

⚠️ Writing to nil map will panic

---

## 3. Read, Write, Delete

```go
m["a"] = 10      // write/update
v := m["a"]      // read
delete(m, "a")   // delete (safe even if missing)
```

---

## 4. Check if Key Exists (VERY IMPORTANT)

```go
v, ok := m["a"]

if ok {
	fmt.Println("exists", v)
} else {
	fmt.Println("not found")
}
```

Why? Because missing keys return zero value.

---

## 5. Iteration

```go
for k, v := range m {
	fmt.Println(k, v)
}
```

⚠️ Order is **random** every time.

---

## 6. Nil vs Empty Map

| Type      | Read | Write   | == nil |
| --------- | ---- | ------- | ------ |
| nil map   | ✅    | ❌ panic | true   |
| empty map | ✅    | ✅       | false  |

```go
var m map[string]int
fmt.Println(m == nil) // true
```

---

## 7. Maps are Reference Types

```go
m1 := map[string]int{"a": 1}
m2 := m1

m2["a"] = 100
fmt.Println(m1["a"]) // 100
```

👉 Both point to same data.

---

## 8. Map Internals (Easy Version)

Think of map like:

```
key → hash → bucket → value
```

* Key is hashed
* Stored in a bucket
* Multiple keys can share a bucket (collision)

---

## 9. Performance Tips

### Preallocate if size known

```go
m := make(map[string]int, 100)
```

### Avoid frequent resizing

---

## 10. Advanced Topics (High-Level)

### A. Collisions

Multiple keys → same bucket → handled internally

### B. Growth

Map grows automatically when needed

### C. Memory

Deleting key doesn’t always shrink map immediately

---

## 11. Concurrency (CRITICAL)

Maps are NOT thread-safe.

```go
// ❌ unsafe
m["a"] = 1
```

Use sync.Mutex or sync.Map:

```go
var m sync.Map
m.Store("a", 1)
```

---

## 12. Common Pitfalls

* Writing to nil map
* Assuming order in iteration
* Concurrent access without lock
* Using non-comparable types as keys (like slices)

```go
// ❌ invalid
map[[]int]string
```

---

## 13. Practice Snippets

### Frequency Counter

```go
arr := []int{1,2,2,3}
freq := make(map[int]int)

for _, v := range arr {
	freq[v]++
}
```

---

### Check Duplicate

```go
seen := make(map[int]bool)
for _, v := range arr {
	if seen[v] {
		fmt.Println("duplicate", v)
	}
	seen[v] = true
}
```

---

## 14. Cheat Sheet

```go
// create
m := make(map[string]int)

// insert
m["a"] = 1

// read
v := m["a"]

// check
v, ok := m["a"]

// delete
delete(m, "a")

// iterate
for k, v := range m {}
```

---

## Final Mental Model

```
Map = reference → hash table → buckets → values
```

---

## One-line takeaway

👉 "Maps store key-value pairs using hashing, and they share memory when copied."
