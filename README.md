<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="assets/go-logo-dark.png">
    <img src="assets/go-logo-light.png" alt="Go logo" width="164">
  </picture>
</p>

<p align="center">
  A structured collection of beginner-friendly Go notes, code examples,<br>
  and practice programs documenting my journey of learning Go from scratch.
</p>

## About Go

Go, often referred to as Golang, is an open-source, statically typed, compiled programming language developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It was announced in 2009, with Go 1.0 officially released in 2012.

Go was designed to simplify the development of large, reliable, and efficient software. It combines the performance and safety of a compiled language with the simplicity and readability of a modern programming language. With built-in support for concurrency, a rich standard library, and fast compilation, Go has become a popular choice for building scalable software.

## Key Features

- **Simple and readable.** Go has a small set of keywords and a consistent coding style, making code written by different developers easy to understand.
- **Fast compilation.** Go compiles programs quickly into a single native binary without requiring an external runtime.
- **Static typing.** Types are checked at compile time, helping catch errors before the program runs.
- **Built-in concurrency.** Goroutines provide lightweight concurrent execution, while channels enable safe communication between them.
- **Garbage collection.** Automatic memory management helps developers focus on writing code instead of manually managing memory.
- **Rich standard library.** Go includes built-in packages for networking, HTTP servers, JSON handling, file operations, testing, cryptography, and more.
- **Cross-platform support.** Go programs can be compiled and run across Windows, macOS, Linux, and other supported platforms.
- **Built-in tooling.** Go provides official tools for formatting, testing, dependency management, documentation, and project management.

## Where Go Is Used

Go is widely used for building fast, reliable, and scalable software, especially in areas where performance and concurrency are important.

- Backend services and REST APIs
- Cloud-native and distributed systems
- DevOps and infrastructure tools
- Command-line (CLI) applications
- Networking software and microservices

Well-known projects built with Go include Docker, Kubernetes, Terraform, and Prometheus.

## Getting Started

Download and install Go from [go.dev/dl](https://go.dev/dl/), then check the installation:

```bash
go version
```

A minimal Go program:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

Save it as `main.go` and run it:

```bash
go run main.go
```

## Common Go Commands

| Command | Description |
| --- | --- |
| `go run file.go` | Compile and run a program |
| `go build` | Compile the package into an executable |
| `go mod init name` | Create a new module |
| `go mod tidy` | Add missing and remove unused dependencies |
| `go fmt ./...` | Format code in the standard Go style |
| `go vet ./...` | Report suspicious code |
| `go test ./...` | Run tests |
| `go doc fmt.Println` | Show documentation for a package or function |

## Learning Resources

- [A Tour of Go](https://go.dev/tour/): interactive introduction to the language
- [Go by Example](https://gobyexample.com/): short, practical examples
- [Effective Go](https://go.dev/doc/effective_go): writing clear, idiomatic Go
- [Go Documentation](https://go.dev/doc/): official guides and references
- [Go Standard Library](https://pkg.go.dev/std): documentation for built-in packages
