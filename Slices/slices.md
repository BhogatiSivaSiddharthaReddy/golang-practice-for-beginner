# Mastering Slices in Go (Advanced Guide)

Slices are the **most important and most misunderstood data structure in Go**. Mastering slices means understanding memory, performance, and real-world Go behavior.

---

## Table of Contents

1. What is a Slice?
2. Slice vs Array (Deep Understanding)
3. Slice Internals (CRITICAL)
4. Slice Header Breakdown
5. Declaration & Initialization
6. Length vs Capacity (Deep Dive)
7. Slicing Semantics (with formulas)
8. Append Internals (VERY IMPORTANT)
9. When Append Reallocates vs Reuses
10. Underlying Array Sharing (BUG SOURCE)
11. Copy vs Reference (Critical Difference)
12. Passing Slices to Functions
13. Nil vs Empty Slice (Real-world impact)
14. Full Slice Expression (Advanced)
15. Memory Leaks with Slices
16. Iteration & Modification Pitfalls
17. Multi-dimensional Slices
18. Performance Optimization Patterns
19. Common Real-world Bugs
20. Hands-on Exercises (Advanced)
21. Cheat Sheet

---

## 1. What is a Slice?

A slice is a **lightweight descriptor that references an underlying array**.

```go
s := []int{1,2,3}
```

👉 It does NOT store data itself.

---

## 2. Slice vs Array (Deep Understanding)

| Feature     | Array     | Slice                 |
| ----------- | --------- | --------------------- |
| Ownership   | Owns data | References data       |
| Copy        | Deep copy | Shallow (header copy) |
| Flexibility | Fixed     | Dynamic               |

👉 Slice is a **view over an array**.

---

## 3. Slice Internals (CRITICAL)

A slice is internally:

```text
pointer → underlying array
length  → number of elements
capacity → max elements before reallocation
```

---

## 4. Slice Header Breakdown

Conceptually:

```go
struct {
    ptr *T
    len int
    cap int
}
```

👉 Copying a slice copies only this header, not data.

---

## 5. Declaration & Initialization

```go
var s []int              // nil slice
s := []int{1,2,3}       // literal
s := make([]int, 3)     // len=3 cap=3
s := make([]int, 3, 5)  // len=3 cap=5
```

---

## 6. Length vs Capacity (Deep Dive)

```go
s := make([]int, 3, 5)
```

```text
len = 3 → usable elements
cap = 5 → max before reallocation
```

---

## 7. Slicing Semantics

```go
s := arr[low:high]
```

```text
len = high - low
cap = cap(arr) - low
```

---

## 8. Append Internals (VERY IMPORTANT)

```go
s := []int{1,2,3}
s = append(s, 4)
```

👉 Two cases:

### Case 1: Capacity available

* Same array reused

### Case 2: Capacity exceeded

* New array allocated
* Old data copied

---

## 9. When Append Reallocates

```go
s := make([]int, 2, 2)
s = append(s, 3)
```

👉 Capacity exceeded → new array created

---

## 10. Underlying Array Sharing (BUG SOURCE)

```go
s1 := []int{1,2,3}
s2 := s1

s2[0] = 100

fmt.Println(s1) // [100 2 3]
```

👉 Both share same array

---

## 11. Copy vs Reference

```go
// WRONG (shared memory)
s2 := s1

// CORRECT (independent copy)
s2 := make([]int, len(s1))
copy(s2, s1)
```

---

## 12. Passing Slices to Functions

```go
func update(s []int) {
    s[0] = 100
}
```

👉 Changes reflect outside

BUT:

```go
s = append(s, 10)
```

👉 Might NOT reflect (new array created)

---

## 13. Nil vs Empty Slice

```go
var s []int   // nil
s := []int{}  // empty
```

| Property  | nil  | empty |
| --------- | ---- | ----- |
| len       | 0    | 0     |
| cap       | 0    | 0     |
| nil check | true | false |

---

## 14. Full Slice Expression (Advanced)

```go
s := arr[low:high:max]
```

👉 Controls capacity

```text
cap = max - low
```

---

## 15. Memory Leak Scenario

```go
large := make([]int, 1000000)
small := large[:10]
```

👉 small slice still holds reference to large array

👉 memory not freed ❌

---

## 16. Iteration Pitfall

```go
for _, v := range s {
    v = 100 // does NOT modify slice
}
```

Correct:

```go
for i := range s {
    s[i] = 100
}
```

---

## 17. Multi-dimensional Slices

```go
matrix := [][]int{
    {1,2,3},
    {4,5,6},
}
```

👉 Not contiguous like arrays

---

## 18. Performance Optimization

### Preallocate

```go
make([]int, 0, 100)
```

👉 avoids repeated allocations

---

## 19. Common Real-world Bugs

### ❌ Shared memory bug

### ❌ Append not reflected

### ❌ Memory leaks from slicing large arrays

### ❌ Modifying loop variable instead of slice

---

## 20. Hands-on Exercises (Advanced)

1. Demonstrate append reallocation
2. Create slice memory leak example
3. Fix shared memory bug
4. Implement dynamic array
5. Write custom append logic

---

## 21. Cheat Sheet

```go
// create
s := []int{1,2,3}

// make
s := make([]int, 3, 5)

// append
s = append(s, 4)

// slice
s2 := s[1:3]

// copy
copy(dest, src)
```

---

## Final Mental Model

```text
Array = actual data
Slice = window over data
Append = may move window to new data
```

---

## One-line takeaway

👉 "Slices are references to arrays, and append can silently change where they point."
