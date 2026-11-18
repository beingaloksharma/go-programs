# Go Programs

A repository of Go programming examples, algorithms, data structures, and foundational exercises.

---

## Repository Structure

```text
go-programs/
├── basics-programs/
│   └── 01_basic_programs/
│       ├── armstrong_number.go     # Armstrong number check using digit extraction and power calculation
│       ├── count_digits.go         # Count digits in an integer using iterative and recursive approaches
│       ├── even_odd.go             # Even or odd number check using modulo arithmetic
│       ├── factorial_recursion.go  # Factorial calculation using recursive and iterative approaches
│       ├── fibonacci_series.go     # Fibonacci series generation using iterative switch-case approach
│       ├── largest_smallest.go     # Find largest and smallest numbers in an integer slice
│       ├── prime_number.go         # Prime number check and range counting using trial division
│       ├── reverse_integer.go      # Integer reversal and palindrome check using arithmetic operations
│       ├── reverse_string.go       # String reversal (backward iteration & two-pointer swap)
│       ├── string_palindrome.go    # String palindrome check using two-pointer reversal
│       ├── sum_of_digits.go        # Sum of digits calculation using iterative and recursive approaches
│       └── swap_numbers.go         # Swap two numbers using tuple assignment, arithmetic, and XOR
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
| [`factorial_recursion.go`](basics-programs/01_basic_programs/factorial_recursion.go) | Computes factorial using recursive and iterative methods | Recursion, Iteration, Base cases, Call stack |
| [`armstrong_number.go`](basics-programs/01_basic_programs/armstrong_number.go) | Checks whether a number is an Armstrong number | Armstrong number, Modulo arithmetic, Digit extraction, Power calculation |
| [`even_odd.go`](basics-programs/01_basic_programs/even_odd.go) | Checks whether a number is even or odd | Modulo arithmetic, Conditional checks, Basic math |
| [`largest_smallest.go`](basics-programs/01_basic_programs/largest_smallest.go) | Finds the smallest and largest numbers in a slice | Array traversal, Min/Max search, Edge case handling |
| [`swap_numbers.go`](basics-programs/01_basic_programs/swap_numbers.go) | Swaps two numbers using multiple techniques | Tuple assignment, Arithmetic operations, Bitwise XOR |
| [`sum_of_digits.go`](basics-programs/01_basic_programs/sum_of_digits.go) | Calculates the sum of digits of an integer | Modulo arithmetic, Digit extraction, Recursion vs Iteration |
| [`count_digits.go`](basics-programs/01_basic_programs/count_digits.go) | Counts the number of digits in an integer | Integer division, Logarithmic complexity, Recursion vs Iteration |

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

# Run factorial calculation
go run basics-programs/01_basic_programs/factorial_recursion.go

# Run armstrong number check
go run basics-programs/01_basic_programs/armstrong_number.go

# Run even or odd check
go run basics-programs/01_basic_programs/even_odd.go

# Run find largest and smallest
go run basics-programs/01_basic_programs/largest_smallest.go

# Run swap numbers
go run basics-programs/01_basic_programs/swap_numbers.go

# Run sum of digits
go run basics-programs/01_basic_programs/sum_of_digits.go

# Run count digits
go run basics-programs/01_basic_programs/count_digits.go
```
