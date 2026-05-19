# Mastering Goroutines in Go

Goroutines are one of the most powerful features in Go.

If you understand this file deeply and practice the examples, you'll understand:
- concurrency
- goroutines
- scheduling
- synchronization
- race conditions
- Go runtime behavior

Goroutines are heavily used in:
- HTTP servers
- Kubernetes
- distributed systems
- microservices
- background workers
- streaming systems

---

# Table of Contents

1. What Goroutines Are
2. Why Goroutines Exist
3. Creating Goroutines
4. Concurrency vs Parallelism
5. Goroutine Scheduling
6. Anonymous Goroutines
7. Goroutines and Functions
8. Goroutines and Closures
9. Synchronization Problem
10. WaitGroup
11. Race Conditions
12. Mutex
13. Channels Introduction
14. Buffered vs Unbuffered Channels
15. Select Statement
16. Goroutine Leaks
17. Common Pitfalls
18. Hands-on Exercises
19. Quick Cheat Sheet

---

# What Goroutines Are

A goroutine is:

```text
a lightweight concurrent function managed by Go runtime
```

---

# Basic Syntax

```go
go function()
```

---

# Example

```go
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go hello()

	time.Sleep(time.Second)
}
```

---

# Important

Without:

```go
time.Sleep()
```

program may exit before goroutine runs.

---

# Why Goroutines Exist

Goroutines allow programs to:
- handle multiple tasks simultaneously
- improve responsiveness
- efficiently utilize CPU resources

---

# Real-world Examples

- HTTP request handling
- background workers
- async logging
- polling systems
- concurrent APIs

---

# Goroutines Are Lightweight

OS threads are expensive.

Goroutines are:
- lightweight
- cheap
- managed by Go runtime

You can create:
- thousands
- even millions

of goroutines.

---

# Concurrency vs Parallelism

---

# Concurrency

```text
multiple tasks progressing independently
```

---

# Parallelism

```text
multiple tasks executing literally at same time
```

---

# Important

Go provides:
- concurrency by design
- parallelism when multiple CPU cores available

---

# Goroutine Scheduling

Go runtime contains scheduler.

Scheduler decides:
- when goroutines run
- where they run
- how they pause/resume

---

# Mental Model

```text
goroutines are multiplexed onto OS threads
```

---

# Simplified Visualization

```text
Many goroutines
       ↓
Go Scheduler
       ↓
Few OS Threads
       ↓
CPU
```

---

# Anonymous Goroutines

---

# Example

```go
go func() {
	fmt.Println("Anonymous goroutine")
}()
```

---

# Why Useful?

Used for:
- inline async work
- callbacks
- background tasks

---

# Goroutines and Functions

---

# Example

```go
func worker(id int) {
	fmt.Println("Worker:", id)
}

func main() {
	go worker(1)

	time.Sleep(time.Second)
}
```

---

# Goroutines and Closures

Very important concept.

---

# Incorrect Example

```go
for i := 0; i < 3; i++ {
	go func() {
		fmt.Println(i)
	}()
}
```

Output unpredictable.

---

# Why?

Closure captures SAME variable.

---

# Correct Version

```go
for i := 0; i < 3; i++ {
	go func(i int) {
		fmt.Println(i)
	}(i)
}
```

---

# Synchronization Problem

Main goroutine may exit before child goroutines finish.

---

# Example

```go
go hello()

fmt.Println("main exits")
```

Possible output:

```text
main exits
```

---

# WaitGroup

Used to wait for goroutines.

---

# Example

```go
package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("working")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go worker(&wg)

	wg.Wait()
}
```

---

# WaitGroup Methods

| Method | Meaning |
|---|---|
| Add(n) | add goroutines |
| Done() | mark complete |
| Wait() | block until complete |

---

# Race Conditions

Race condition happens when:
- multiple goroutines
- access shared data
- simultaneously
- without synchronization

---

# Example

```go
var counter int

func increment() {
	counter++
}
```

Multiple goroutines running this causes race condition.

---

# Why Dangerous?

Because operations are NOT atomic.

Example:

```text
read → modify → write
```

can interleave unpredictably.

---

# Mutex

Mutex protects shared data.

---

# Example

```go
var mu sync.Mutex
var counter int

func increment() {
	mu.Lock()

	counter++

	mu.Unlock()
}
```

---

# Mental Model

```text
mutex = only one goroutine allowed inside critical section
```

---

# Channels Introduction

Channels allow goroutines to communicate safely.

---

# Creating Channel

```go
ch := make(chan int)
```

---

# Sending

```go
ch <- 10
```

---

# Receiving

```go
value := <-ch
```

---

# Example

```go
package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	value := <-ch

	fmt.Println(value)
}
```

---

# Important

Channels synchronize goroutines.

---

# Buffered Channels

---

# Example

```go
ch := make(chan int, 2)
```

Capacity:
- 2 values

---

# Unbuffered Channels

```go
ch := make(chan int)
```

Send blocks until receive happens.

---

# Buffered vs Unbuffered

| Type | Behavior |
|---|---|
| unbuffered | sender waits |
| buffered | sender waits only when full |

---

# Select Statement

Used to wait on multiple channels.

---

# Example

```go
select {
case msg := <-ch1:
	fmt.Println(msg)

case msg := <-ch2:
	fmt.Println(msg)
}
```

---

# Why Useful?

Used for:
- multiplexing
- timeouts
- cancellation
- async coordination

---

# Timeout Example

```go
select {
case msg := <-ch:
	fmt.Println(msg)

case <-time.After(time.Second):
	fmt.Println("timeout")
}
```

---

# Goroutine Leaks

Happens when goroutine:
- blocks forever
- never exits

---

# Example

```go
go func() {
	<-ch
}()
```

If nobody sends:
- goroutine leaks

---

# Important Rule

Always ensure:
- goroutines can exit
- channels eventually close
- receivers/senders match

---

# Common Pitfalls

---

# 1. Forgetting synchronization

```go
go hello()
```

Program exits early.

---

# 2. Race conditions

Shared data without mutex/channel.

---

# 3. Closure capture bug

Loop variable captured incorrectly.

---

# 4. Deadlocks

All goroutines waiting forever.

---

# Example

```go
ch := make(chan int)

ch <- 10
```

No receiver → deadlock.

---

# 5. Goroutine leaks

Blocked forever on channel.

---

# Hands-on Exercises

1. Create basic goroutine
2. Create anonymous goroutine
3. Pass parameters to goroutine
4. Use WaitGroup
5. Create race condition
6. Fix using Mutex
7. Create channels
8. Use select statement
9. Create timeout example

---

# Quick Cheat Sheet

```go
// start goroutine
go hello()

// anonymous goroutine
go func() {}()

// waitgroup
var wg sync.WaitGroup

wg.Add(1)
wg.Done()
wg.Wait()

// mutex
mu.Lock()
mu.Unlock()

// channel
ch := make(chan int)

// send
ch <- 10

// receive
x := <-ch

// buffered channel
ch := make(chan int, 2)

// select
select {}
```

---

# Final Mental Model

```text
Goroutines enable lightweight concurrency.
```

Go concurrency is built around:
- goroutines
- channels
- synchronization primitives

---

# Deep Insight

```text
Concurrency is about coordination,
not just parallel execution.
```

---

# One-line Takeaway

👉 “Goroutines are lightweight concurrent functions managed efficiently by the Go runtime.”