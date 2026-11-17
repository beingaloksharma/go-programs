# Go Programs

A repository of Go programming examples, algorithms, data structures, and foundational exercises.

---

## Repository Structure

```text
go-programs/
├── basics-programs/
│   └── 01_basic_programs/
│       ├── fibonacci_series.go     # Fibonacci series generation using iterative switch-case approach
│       ├── prime_number.go         # Prime number check and range counting using trial division
│       ├── reverse_integer.go      # Integer reversal and palindrome check using arithmetic operations
│       ├── reverse_string.go       # String reversal (backward iteration & two-pointer swap)
│       └── string_palindrome.go    # String palindrome check using two-pointer reversal
├── go.mod
└── README.md
```

---

## Programs Catalog

### 01. Basic Programs

| Program | Description | Key Concepts |
|---|---|---|
| [`reverse_string.go`](basics-programs/01_basic_programs/reverse_string.go) | Reverses strings using multiple approaches | Runes, Unicode handling, Two-pointer technique, Slice pre-allocation |
| [`string_palindrome.go`](basics-programs/01_basic_programs/string_palindrome.go) | Checks whether a string is a palindrome | Palindrome check, Runes, Two-pointer technique |
| [`reverse_integer.go`](basics-programs/01_basic_programs/reverse_integer.go) | Reverses an integer and checks if it is a palindrome | Modulo arithmetic, Palindrome check, Digit extraction |
| [`prime_number.go`](basics-programs/01_basic_programs/prime_number.go) | Checks if a number is prime and counts primes in a range | Prime numbers, Trial division up to sqrt(n), Optimization |
| [`fibonacci_series.go`](basics-programs/01_basic_programs/fibonacci_series.go) | Generates Fibonacci numbers for a given count | Fibonacci sequence, Switch-case, Iterative calculation |

---

## Running Programs

```bash
# Run reverse string
go run basics-programs/01_basic_programs/reverse_string.go

# Run string palindrome check
go run basics-programs/01_basic_programs/string_palindrome.go

# Run reverse integer
go run basics-programs/01_basic_programs/reverse_integer.go

# Run prime number check
go run basics-programs/01_basic_programs/prime_number.go

# Run fibonacci series
go run basics-programs/01_basic_programs/fibonacci_series.go
```
