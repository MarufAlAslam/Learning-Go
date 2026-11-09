# Learning-Go

A 60-day roadmap from zero to advanced Go. Aim for 1–2 hours per day: read, code along, then do the exercise. Each day gets its own folder (`day-1`, `day-2`, ...).

## Phase 1: Foundations (Days 1–10)

- [ ] **[Day 1](day-1/README.md)** – Install Go, set up the editor, `go run`, `go build`, Hello World, how a Go program is structured (`package main`, `func main`)
- [ ] **[Day 2](day-2/README.md)** – Variables, constants, basic types, zero values, type conversion, `fmt` printing
- [ ] **[Day 3](day-3/README.md)** – Operators, `if/else`, `switch`, user input with `fmt.Scan`/`bufio`
- [ ] **[Day 4](day-4/README.md)** – Loops: `for` (all forms), `break`, `continue`, labels, `range`
- [ ] **[Day 5](day-5/README.md)** – Functions: parameters, multiple returns, named returns, variadic functions
- [ ] **[Day 6](day-6/README.md)** – Functions deep-dive: first-class functions, closures, anonymous functions, `defer`
- [ ] **[Day 7](day-7/README.md)** – Arrays and slices: length, capacity, `append`, `copy`, slicing
- [ ] **[Day 8](day-8/README.md)** – Maps: create, read, delete, check existence, iterate; slices of maps
- [ ] **[Day 9](day-9/README.md)** – Strings, runes, bytes, `strings`, `strconv`, `unicode` packages
- [ ] **[Day 10](day-10/README.md)** – Mini project: CLI calculator or number-guessing game. Review of Phase 1

## Phase 2: Core Language (Days 11–20)

- [ ] **[Day 11](day-11/README.md)** – Pointers: `&`, `*`, pass by value vs reference, `new`
- [ ] **[Day 12](day-12/README.md)** – Structs: fields, methods, embedding, struct tags
- [ ] **[Day 13](day-13/README.md)** – Methods: value vs pointer receivers, method sets
- [ ] **[Day 14](day-14/README.md)** – Interfaces: implicit implementation, empty interface, `any`
- [ ] **[Day 15](day-15/README.md)** – Interfaces deep-dive: type assertions, type switches, `Stringer`, common stdlib interfaces
- [ ] **[Day 16](day-16/README.md)** – Error handling: `error` interface, custom errors, `errors.New`, `fmt.Errorf` with `%w`
- [ ] **[Day 17](day-17/README.md)** – Error handling deep-dive: `errors.Is`, `errors.As`, `panic`, `recover`
- [ ] **[Day 18](day-18/README.md)** – Packages and modules: `go mod init`, imports, exported vs unexported, `go get`, `go mod tidy`
- [ ] **[Day 19](day-19/README.md)** – Project layout: multi-package projects, `internal/`, `cmd/`, init functions
- [ ] **[Day 20](day-20/README.md)** – Mini project: to-do list CLI app with structs, slices, and JSON file storage

## Phase 3: Intermediate Go (Days 21–32)

- [ ] **[Day 21](day-21/README.md)** – Generics: type parameters, constraints, generic functions
- [ ] **[Day 22](day-22/README.md)** – Generics deep-dive: generic types, `comparable`, `cmp`, `slices` and `maps` packages
- [ ] **[Day 23](day-23/README.md)** – Standard library tour: `os`, `io`, `bufio`, reading and writing files
- [ ] **[Day 24](day-24/README.md)** – `encoding/json`: marshal, unmarshal, struct tags, custom marshalers
- [ ] **[Day 25](day-25/README.md)** – `time` package: durations, formatting, timers, tickers
- [ ] **[Day 26](day-26/README.md)** – `flag`, environment variables, `log` and `log/slog` (structured logging)
- [ ] **[Day 27](day-27/README.md)** – Testing basics: `go test`, table-driven tests, `t.Run`, `t.Helper`
- [ ] **[Day 28](day-28/README.md)** – Testing deep-dive: coverage, benchmarks, examples, fuzzing, `testify` (optional)
- [ ] **[Day 29](day-29/README.md)** – Mocking and testing with interfaces, `httptest`, test fixtures
- [ ] **[Day 30](day-30/README.md)** – Tooling: `go fmt`, `go vet`, `golangci-lint`, `go doc`, `go generate`
- [ ] **[Day 31](day-31/README.md)** – Regular expressions, `text/template`, `html/template`
- [ ] **[Day 32](day-32/README.md)** – Mini project: CLI tool (e.g., file organizer or log parser) with tests

## Phase 4: Concurrency (Days 33–42)

- [ ] **[Day 33](day-33/README.md)** – Goroutines: concepts, `go` keyword, the scheduler (GMP model)
- [ ] **[Day 34](day-34/README.md)** – `sync.WaitGroup`, race conditions, the race detector (`go run -race`)
- [ ] **[Day 35](day-35/README.md)** – Channels: unbuffered vs buffered, send/receive, closing, `range` over channels
- [ ] **[Day 36](day-36/README.md)** – `select`: multiplexing, timeouts, default case, non-blocking operations
- [ ] **[Day 37](day-37/README.md)** – `sync.Mutex`, `RWMutex`, `sync.Once`, `sync/atomic`
- [ ] **[Day 38](day-38/README.md)** – `context`: cancellation, deadlines, timeouts, passing values
- [ ] **[Day 39](day-39/README.md)** – Concurrency patterns: worker pool, fan-out/fan-in, pipeline
- [ ] **[Day 40](day-40/README.md)** – More patterns: rate limiting, semaphores, `errgroup`, `sync.Pool`, `sync.Map`
- [ ] **[Day 41](day-41/README.md)** – Common concurrency bugs: deadlocks, goroutine leaks, data races, and how to debug them
- [ ] **[Day 42](day-42/README.md)** – Mini project: concurrent web crawler or parallel file processor

## Phase 5: Web and Databases (Days 43–52)

- [ ] **[Day 43](day-43/README.md)** – `net/http`: HTTP client and server, handlers, `ServeMux` (Go 1.22+ routing)
- [ ] **[Day 44](day-44/README.md)** – Building a REST API: routing, JSON request/response, status codes
- [ ] **[Day 45](day-45/README.md)** – Middleware: logging, recovery, auth, CORS; request context
- [ ] **[Day 46](day-46/README.md)** – Routers and frameworks: `chi` or `gin`; validation; project structure
- [ ] **[Day 47](day-47/README.md)** – SQL with `database/sql`: connections, queries, prepared statements, transactions
- [ ] **[Day 48](day-48/README.md)** – PostgreSQL with `pgx`, migrations (`golang-migrate`), `sqlc` or GORM
- [ ] **[Day 49](day-49/README.md)** – Authentication: password hashing (bcrypt), JWT, sessions
- [ ] **[Day 50](day-50/README.md)** – Configuration, graceful shutdown, Docker, and Dockerfile for Go
- [ ] **[Day 51](day-51/README.md)** – Caching with Redis, background jobs, and queues
- [ ] **[Day 52](day-52/README.md)** – Mini project: full REST API (users, auth, CRUD) with Postgres and tests

## Phase 6: Advanced Go (Days 53–60)

- [ ] **[Day 53](day-53/README.md)** – gRPC and Protocol Buffers
- [ ] **[Day 54](day-54/README.md)** – Reflection (`reflect`), `unsafe`, and when not to use them
- [ ] **[Day 55](day-55/README.md)** – Memory model: stack vs heap, escape analysis, garbage collector internals
- [ ] **[Day 56](day-56/README.md)** – Profiling and optimization: `pprof`, `go tool trace`, benchmarks, `benchstat`
- [ ] **[Day 57](day-57/README.md)** – Design patterns in Go, clean/hexagonal architecture, dependency injection
- [ ] **[Day 58](day-58/README.md)** – Observability: metrics (Prometheus), tracing (OpenTelemetry), structured logs
- [ ] **[Day 59](day-59/README.md)** – Cross-compilation, build tags, CI/CD with GitHub Actions, releasing binaries
- [ ] **[Day 60](day-60/README.md)** – Capstone project: production-grade service (API + DB + concurrency + tests + Docker + CI)

## After Day 60

- Read the Go source code of popular libraries (`net/http`, `chi`, `cobra`)
- Contribute to an open-source Go project
- Build: a key-value store, a load balancer, a CLI with `cobra`, a message queue
- Read: *The Go Programming Language* (Donovan & Kernighan), *Concurrency in Go* (Katherine Cox-Buday), *100 Go Mistakes* (Teiva Harsanyi)

## Resources

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Standard Library Docs](https://pkg.go.dev/std)
- [Go Blog](https://go.dev/blog/)
