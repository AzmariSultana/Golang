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

Go, often called Golang, is an open-source, statically typed and compiled programming language designed at Google by Robert Griesemer, Rob Pike and Ken Thompson. It was announced in 2009, and version 1.0 was released in 2012.

Go was created to make building large, reliable software simpler. It combines the speed and safety of a compiled language with the readability of a scripting language, and it has built-in support for concurrency.

## Key Features

- **Simple and readable.** Go has a small set of keywords and one standard code style, so code written by different people looks familiar.
- **Fast compilation.** Even large projects compile in seconds into a single native binary with no external runtime.
- **Static typing.** Types are checked at compile time, catching many errors before the program runs.
- **Concurrency built in.** Goroutines are lightweight threads, and channels let them communicate safely.
- **Garbage collection.** Memory is managed automatically.
- **Rich standard library.** Networking, HTTP servers, JSON, file handling, testing and cryptography are included.
- **Cross-platform.** The same code can be compiled for Windows, macOS and Linux.
- **Built-in tooling.** Formatting, testing, dependency management and documentation tools ship with the language.

## Where Go Is Used

- Backend services and REST APIs
- Cloud-native and distributed systems
- DevOps and infrastructure tools
- Command-line applications
- Networking and microservices

Well-known projects written in Go include Docker, Kubernetes, Terraform and Prometheus.

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
