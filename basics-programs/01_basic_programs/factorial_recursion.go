// Package main provides basic programs to calculate the factorial of a number in Go.
//
// Program: Factorial Calculation
// Methods demonstrated:
//  1. Recursive approach
//  2. Iterative approach
package main

import "fmt"

func main() {
	// Sample integer to test
	num := 5

	// Approach 1: Calculate factorial using recursion
	factRes1 := factRecursive(num)
	fmt.Printf("Factorial of %d is %d (Recursive)\n", num, factRes1)

	// Approach 2: Calculate factorial using iteration
	factRes2 := fact(num)
	fmt.Printf("Factorial of %d is %d (Iterative)\n", num, factRes2)
}

// factRecursive calculates the factorial of a non-negative integer using recursion.
//
// Base cases: 0! = 1, 1! = 1
// Recursive case: n! = n * (n - 1)!
//
// Time Complexity: O(n), Space Complexity: O(n) (due to call stack frames)
func factRecursive(n int) int {
	// Base case: factorial of 0 or 1 is 1
	if n <= 1 {
		return 1
	}

	// Recursive call
	return n * factRecursive(n-1)
}

// fact calculates the factorial of a non-negative integer using iteration.
//
// Time Complexity: O(n), Space Complexity: O(1)
func fact(n int) int {
	// Base case: factorial of 0 or 1 is 1
	if n == 0 || n == 1 {
		return 1
	}

	// Accumulator for factorial product
	f := 1

	// Multiply numbers from n down to 1
	for n > 0 {
		f = f * n
		n = n - 1
	}

	return f
}
