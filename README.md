# Golang

My notes and practice code while learning Go (Golang) from the basics.

## About Go

Go is an open-source programming language created at Google by Robert Griesemer, Rob Pike and Ken Thompson, and released in 2009. It was designed to be simple, fast and reliable for building large software.

Some things that make Go stand out:

- **Simple syntax.** The language has a small set of keywords, so it is quick to learn and easy to read.
- **Compiled and fast.** Go compiles to a single native binary with no extra runtime to install.
- **Statically typed.** Types are checked at compile time, which catches many mistakes before the program runs.
- **Built-in concurrency.** Goroutines and channels make it easy to run many tasks at the same time.
- **Garbage collected.** Memory is managed automatically.
- **Strong standard library.** Packages for HTTP servers, JSON, file handling, testing and more come with Go.
- **Good tooling.** `go fmt`, `go test`, `go vet` and `go mod` are included out of the box.

Go is widely used for backend services, APIs, cloud and DevOps tools, and command-line programs. Docker, Kubernetes and Terraform are all written in Go.

## Getting Started

1. Download and install Go from [go.dev/dl](https://go.dev/dl/).
2. Check the installation:

   ```bash
   go version
   ```

3. Run a single file:

   ```bash
   go run values.go
   ```

A minimal Go program looks like this:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Repository Structure

| Folder | Topics |
| --- | --- |
| [Fundamentals](Fundamentals) | Values and printing, variables, constants, enums with `iota`, plus a PDF of study notes |

### Fundamentals

| File | Topic |
| --- | --- |
| `values.go` | Basic data types and printing with `fmt` |
| `variable.go` | Declaring variables, zero values and `:=` |
| `constants.go` | Constants, scope, `float32` vs `float64` |
| `enums.go` | Enums using `const` blocks and `iota` |
| `Go_Fundamentals.pdf` | Beginner study notes for all four topics |

Each file has its own `main` function, so run them one at a time, for example `go run Fundamentals/enums.go`.

## Useful Resources

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Documentation](https://go.dev/doc/)
