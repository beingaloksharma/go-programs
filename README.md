# Go Programs

A repository of Go programming examples, algorithms, data structures, and foundational exercises.

---

## Repository Structure

```text
go-programs/
├── basics-programs/
│   └── 01_basic_programs/
│       ├── reverse_string.go       # Implementations of string reversal (backward iteration & two-pointer)
│       └── reverse_string_test.go  # Unit tests and benchmarks
├── go.mod
└── README.md
```

---

## Programs Catalog

### 01. Basic Programs

| Program | Description | Key Concepts |
|---|---|---|
| [`reverse_string.go`](basics-programs/01_basic_programs/reverse_string.go) | Reverses strings using multiple approaches | Runes, Unicode handling, Two-pointer technique, Slice allocation |

---

## Running & Testing

### Run a Program
```bash
go run basics-programs/01_basic_programs/reverse_string.go
```

### Run Unit Tests
```bash
go test -v ./...
```

### Run Benchmarks
```bash
go test -bench=. -benchmem ./...
```
