# Go Basics: Packages, Modules, Imports, and Core Commands

This repository is a starting point for anyone who wants to learn the Go programming language.

This README summarizes the Go concepts we learned:

- how Go manages dependencies
- the difference between `go get` and `go install`
- what `go.mod` and `go.sum` do
- the importance of the `package` keyword
- the importance of the `import` keyword
- the difference between `package main` and other packages
- basic Go commands such as `go run`, `go mod init`, and `go build`

---

## 1. Does Go Have a Package Manager?

Yes, Go has dependency management built into the `go` command itself.

Unlike Python, where we usually use `pip`, Go uses:

- `go mod`
- `go get`
- `go install`
- `go mod tidy`

So Go does not usually need a separate package manager like `pip` or `npm`.
Its module system is built into the language tooling.

---

## 2. `go get` vs `go install`

These two commands are related, but they are used for different purposes.

### `go get`

`go get` is mainly used to add or update a dependency in your project.

Example:

```bash
go get github.com/gin-gonic/gin
```

This updates your project's dependency information in `go.mod`.

Use `go get` when:

- you want to use an external library in your project
- you want to add or upgrade a dependency

### `go install`

`go install` is mainly used to install a Go executable tool.

Example:

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

This installs a command-line tool, not a library for your current project.

Use `go install` when:

- you want to install a Go-based CLI tool
- you want a binary executable available on your system

### Quick Comparison

- `go get` adds or updates a project dependency
- `go install` installs an executable tool

---

## 3. `go.mod` vs `go.sum`

These two files work together.

### `go.mod`

`go.mod` describes your module and its direct dependencies.

It contains:

- module name
- Go version
- required dependencies

Example:

```go
module myapp

go 1.24

require github.com/gin-gonic/gin v1.10.0
```

Think of `go.mod` as:

**"What my project needs."**

### `go.sum`

`go.sum` stores checksums (hashes) of dependency versions.

It is used to:

- verify downloaded dependencies
- make builds more reproducible
- improve security and consistency

Think of `go.sum` as:

**"Proof that the downloaded dependency files are exactly what Go expects."**

### Simple Analogy

- `go.mod` = shopping list
- `go.sum` = receipt + verification record

---

## 4. Importance of the `package` Keyword

In Go, every `.go` file must begin with a `package` declaration.

Example:

```go
package main
```

or

```go
package mathutil
```

The `package` keyword tells Go:

- which package this file belongs to
- how code is grouped
- how functions, variables, and types are organized

### Why it is important

The `package` keyword is important because it helps:

- organize code into logical units
- separate reusable code from executable code
- control visibility and structure
- make code easier to maintain

Without packages, large Go programs would become difficult to manage.

---

## 5. Importance of the `import` Keyword

The `import` keyword is used to bring code from another package into the current file.

Example:

```go
import "fmt"
```

This allows us to use functions from the `fmt` package, such as:

```go
fmt.Println("Hello")
```

### Why `import` is important

The `import` keyword is important because it lets us:

- use Go standard library packages
- use external libraries
- reuse code instead of rewriting it
- keep programs modular and clean

Example with multiple imports:

```go
import (
	"fmt"
	"net/http"
)
```

---

## 6. Difference Between `package main` and Other Packages

This is one of the most important concepts in Go.

### `package main`

`package main` is special in Go.

It tells Go:

**"This package is the entry point of an executable program."**

If a package is named `main`, it can contain a `main()` function, and Go will run that function when the program starts.

Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go")
}
```

This creates a runnable program.

### Other packages

Other package names are usually used for reusable code.

Example:

```go
package mathutil

func Add(a, b int) int {
	return a + b
}
```

This package is not directly executable.
It is meant to be imported and used by another file, often from `package main`.

### Key Difference

- `package main` is used to create executable programs
- other packages are used to create reusable libraries or modules

---

## 7. Basic Go Commands

Here are some important beginner Go commands.

### `go mod init`

This command creates a new Go module in your project.

Example:

```bash
go mod init myapp
```

It creates a `go.mod` file.

Use it when:

- you start a new Go project
- you want to enable Go modules in the project

### `go run`

This command compiles and runs your Go program directly without creating a permanent executable file.

Example:

```bash
go run main.go
```

Use it when:

- you want to quickly test or run your program
- you are still developing and do not need a built binary yet

### `go build`

This command compiles your Go code and creates an executable binary.

Example:

```bash
go build
```

or

```bash
go build -o app
```

Use it when:

- you want to create a runnable binary
- you want to check whether your code compiles successfully

### `go mod tidy`

This command cleans up module dependencies.

It:

- adds missing required modules
- removes unused dependencies
- updates `go.mod` and `go.sum`

Example:

```bash
go mod tidy
```

### Quick Command Summary

- `go mod init` creates a new module
- `go run` compiles and runs code directly
- `go build` builds an executable binary
- `go get` adds or updates dependencies
- `go install` installs executable tools
- `go mod tidy` cleans and syncs dependencies

---

## 8. Example of `package main` Using Another Package

### `main.go`

```go
package main

import (
	"fmt"
	"myapp/mathutil"
)

func main() {
	fmt.Println(mathutil.Add(2, 3))
}
```

### `mathutil/mathutil.go`

```go
package mathutil

func Add(a, b int) int {
	return a + b
}
```

### What happens here

- `mathutil` is a reusable package
- `main` is the executable package
- `main.go` imports `mathutil`
- the program runs from `main()`

---

## 9. Summary

### Go dependency management

- Go has built-in module support
- it does not usually need a separate package manager like `pip`

### Commands

- `go mod init` starts a new module
- `go run` runs the program directly
- `go build` compiles the program into a binary
- `go get` adds or updates a dependency
- `go install` installs an executable tool
- `go mod tidy` cleans and syncs dependencies

### Files

- `go.mod` defines project dependencies
- `go.sum` stores dependency checksums for verification

### Keywords

- `package` defines which package a file belongs to
- `import` brings another package into the file

### Package types

- `package main` is the executable entry point
- other packages are reusable code libraries

---

## 10. Final Note

A good way to think about Go is:

- packages organize code
- imports connect code
- modules manage dependencies
- `main` runs the program

These concepts form the foundation of writing Go applications.
