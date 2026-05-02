
# Mastering File Processing in Go

If you understand this file deeply and practice the examples, you'll have a **strong, practical understanding of file handling in Go**, from basics to real-world usage.

---

## Table of Contents

1. What File Processing Is
2. How Files Work Internally
3. File Processing Lifecycle
4. The Role of `os` Package
5. The Role of `bufio` Package
6. Why `os` and `bufio` Work Together
7. Reading Files (All Approaches)
8. Writing and Appending Files
9. Processing Files (Real Patterns)
10. Streams and Memory Model
11. Error Handling and Safety
12. Performance Considerations
13. Common Pitfalls
14. Real-world Patterns
15. Hands-on Exercises
16. Quick Cheat Sheet

---

## What File Processing Is

File processing means working with data stored on disk.

Typical flow:

```

Open → Read/Write → Process → Close

```

A file is simply a **sequence of bytes**:

```

Hello world
Go is awesome

```

Internally:

```

[72 101 108 108 111 ...]

```

---

## How Files Work Internally

```

Disk → Operating System → Go Runtime → Your Program

````

- Files exist on disk
- OS controls access (permissions, locking, etc.)
- Go interacts via system calls through the `os` package

---

## File Processing Lifecycle

Every file operation follows:

1. Open the file
2. Perform read/write
3. Process the data
4. Close the file

Example structure:

```go
file, err := os.Open("file.txt")
if err != nil {
    return
}
defer file.Close()
````

---

## The Role of `os` Package

The `os` package is responsible for **interacting with the operating system**.

It provides:

* Opening files
* Creating files
* Writing files
* Managing file modes and permissions

Example:

```go
file, err := os.Open("file.txt")
```

Important:

* This does NOT read the file
* It only gives you access (a file handle)

---

## The Role of `bufio` Package

The `bufio` package is used for **efficient reading and writing**.

Without buffering:

```
Each read → system call → slow ❌
```

With buffering:

```
Read large chunk → process in memory → fast ✅
```

Example:

```go
scanner := bufio.NewScanner(file)
```

---

## Why `os` and `bufio` Work Together

```
os     → opens file
bufio  → processes data efficiently
```

Example:

```go
file, _ := os.Open("file.txt")
scanner := bufio.NewScanner(file)
```

Important rule:

* `bufio` cannot open files
* It needs an `io.Reader` (like `os.File`)

---

## Reading Files (All Approaches)

### 1. Read entire file (small files)

```go
data, err := os.ReadFile("file.txt")
if err != nil {
    return
}
fmt.Println(string(data))
```

Use when:

* file is small
* simplicity matters

---

### 2. Read line by line (recommended)

```go
file, err := os.Open("file.txt")
if err != nil {
    return
}
defer file.Close()

scanner := bufio.NewScanner(file)

for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}
```

Use when:

* file is large
* streaming processing needed

---

## Writing and Appending Files

### Writing (overwrite)

```go
os.WriteFile("output.txt", []byte("Hello Go"), 0644)
```

---

### Appending

```go
file, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
defer file.Close()

file.WriteString("New line\n")
```

---

## Processing Files (Real Patterns)

### Word Count

```go
file, _ := os.Open("file.txt")
defer file.Close()

scanner := bufio.NewScanner(file)

count := 0

for scanner.Scan() {
    words := strings.Fields(scanner.Text())
    count += len(words)
}

fmt.Println("Total words:", count)
```

---

### Log Analyzer

```
INFO User logged in
ERROR DB failed
ERROR Timeout
```

```go
file, _ := os.Open("logs.txt")
defer file.Close()

scanner := bufio.NewScanner(file)

errors := 0

for scanner.Scan() {
    line := scanner.Text()

    if strings.Contains(line, "ERROR") {
        errors++
        fmt.Println("Error:", line)
    }
}

fmt.Println("Total errors:", errors)
```

---

## Streams and Memory Model

Files are processed as streams:

```
Read → Process → Discard → Next
```

Key idea:

* You don’t need entire file in memory
* You process piece by piece

---

## Error Handling and Safety

Always handle errors:

```go
if err != nil {
    return
}
```

Always close files:

```go
defer file.Close()
```

Why?

* prevents memory leaks
* avoids file locking issues

---

## Performance Considerations

| Method          | Use Case         |
| --------------- | ---------------- |
| `os.ReadFile`   | small files      |
| `bufio.Scanner` | large files      |
| `bufio.Reader`  | advanced control |

---

## Common Pitfalls

### 1. Writing to nil file

```go
var f *os.File
f.WriteString("hello") // panic
```

---

### 2. Forgetting to close file

Leads to:

* resource leaks
* file locks

---

### 3. Using ReadFile on large files

```go
os.ReadFile("huge.txt") // bad idea
```

---

### 4. Ignoring scanner errors

```go
if err := scanner.Err(); err != nil {
    fmt.Println(err)
}
```

---

## Real-world Patterns

You will use file processing for:

* log parsing
* CSV processing
* config loading
* data pipelines
* CLI tools

---

## Hands-on Exercises

1. Read a file and count lines
2. Count words in a file
3. Filter lines containing a keyword
4. Copy one file to another
5. Append logs to a file
6. Build a mini log analyzer

---

## Quick Cheat Sheet

```go
// open
file, _ := os.Open("file.txt")
defer file.Close()

// read all
data, _ := os.ReadFile("file.txt")

// scanner
scanner := bufio.NewScanner(file)

// write
os.WriteFile("file.txt", []byte("data"), 0644)

// append
file, _ := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 0644)

// process loop
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

---

## Final Mental Model

```
File = stream of bytes
os   = access to file
bufio = efficient processing
```

---

## One-line takeaway

 “File processing in Go is about streaming data from disk, processing it efficiently, and managing resources safely.”
