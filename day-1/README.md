# Day 1: Setup and Hello World

[← Back to roadmap](../README.md)

## Status check

- [x] Go installed (`go1.27.3 darwin/arm64`)
- [x] `hello.go` created and runs correctly

## Things to learn

1. **What Go is**: compiled, statically typed, fast, built-in concurrency, single binary output
2. **Go toolchain commands**
   - `go version`: check the installed version
   - `go run file.go`: compile and run in one step
   - `go build file.go`: produce an executable
   - `go fmt` / `gofmt`: auto-format code
   - `go env`: show Go environment (`GOPATH`, `GOROOT`)
3. **Anatomy of a Go program**
   - `package main`: marks an executable program
   - `import`: bringing in packages (single and grouped imports)
   - `func main()`: the entry point
4. **The `fmt` package**: `Println`, `Print`, `Printf`, and basic verbs (`%s`, `%d`, `%v`)
5. **Basics of syntax**: comments (`//`, `/* */`), no semicolons, braces must be on the same line, unused imports/variables are compile errors

## Files to create

```
day-1/
├── README.md        (this file)
├── hello.go         (done) | Println "Hello, World!"
├── greet.go         (done) | Printf with name, age, and %s / %d / %v
├── comments.go      (done) | single-line and block comments, a documented function
├── imports.go       (done) | grouped imports (fmt, os, runtime) printing OS and Go version
└── exercise.go      (done) | your own introduction program (see below)
```

Each file is its own `package main` with its own `main()`, so run them separately (`go run greet.go`). Don't use `go run .` in this folder, since multiple `main` functions in one package will conflict.

## Tasks

1. Run `hello.go` with `go run hello.go`.
2. Build it with `go build hello.go`, then run the binary (`./hello`). Note the file size and that it needs no Go installed to run.
3. **greet.go**: store your name and age in variables and print them with `Printf`.
4. **comments.go**: add both comment styles and a doc comment above a function.
5. **imports.go**: print `runtime.GOOS`, `runtime.GOARCH`, and `runtime.Version()`.
6. **Break things on purpose**: add an unused import, remove the `package` line, and misspell `Println`. Read each compiler error.
7. Run `go fmt ./...` after messing up the indentation, and see it fix the file.

## Exercise

Write `exercise.go` that prints:

```
Name: <your name>
Goal: Learn Go in 60 days
Today is day 1
```

Use `Printf` with variables, not hardcoded strings.

## Questions to answer (in your own words)

- What is the difference between `go run` and `go build`?
- Why does the program need `package main` and `func main()`?
- What happens if you import a package and don't use it?

## Resources

- [A Tour of Go: Basics](https://go.dev/tour/basics/1)
- [Go by Example: Hello World](https://gobyexample.com/hello-world)
- [fmt package docs](https://pkg.go.dev/fmt)
