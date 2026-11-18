// Package main provides a basic program to calculate the sum of digits of a number in Go.
//
// Program: Sum of Digits
// Methods demonstrated:
//  1. Iterative approach using modulo arithmetic
//  2. Recursive approach
package main

import "fmt"

func main() {
	// Sample number to test
	digits := 123

	// Approach 1: Iterative calculation
	fmt.Printf("Sum of Digits of %d is %d (Iterative)\n", digits, sumOfDigits(digits))

	// Approach 2: Recursive calculation
	fmt.Printf("Sum of Digits of %d is %d (Recursive)\n", digits, sumOfDigitsRecursive(digits))
}

// sumOfDigits calculates the sum of all digits in an integer using iteration.
//
// Time Complexity: O(log10(n)), Space Complexity: O(1)
func sumOfDigits(digits int) int {
	// Handle negative numbers by converting to positive
	if digits < 0 {
		digits = -digits
	}

	var total int
	// Extract digits from right to left using modulo and division
	for digits > 0 {
		digit := digits % 10
		total += digit
		digits = digits / 10
	}
	return total
}

// sumOfDigitsRecursive calculates the sum of all digits in an integer using recursion.
//
// Time Complexity: O(log10(n)), Space Complexity: O(log10(n)) (call stack)
func sumOfDigitsRecursive(digits int) int {
	// Handle negative numbers
	if digits < 0 {
		digits = -digits
	}

	// Base case: single-digit number
	if digits < 10 {
		return digits
	}

	// Recursive step: last digit + sum of remaining digits
	return (digits % 10) + sumOfDigitsRecursive(digits/10)
}
