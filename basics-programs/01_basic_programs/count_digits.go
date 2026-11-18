// Package main provides a basic program to count the number of digits in an integer in Go.
//
// Program: Count Digits
// Methods demonstrated:
//  1. Iterative approach using division
//  2. Recursive approach
package main

import "fmt"

func main() {
	// Sample integer to test
	digit := 123456

	// Approach 1: Iterative counting
	fmt.Printf("Count is %d for digits %d (Iterative)\n", countDigits(digit), digit)

	// Approach 2: Recursive counting
	fmt.Printf("Count is %d for digits %d (Recursive)\n", countDigitsRecursive(digit), digit)
}

// countDigits counts the total number of digits in an integer using iteration.
//
// Time Complexity: O(log10(n)), Space Complexity: O(1)
func countDigits(num int) int {
	// 0 has exactly 1 digit
	if num == 0 {
		return 1
	}

	// Handle negative numbers by converting to positive
	if num < 0 {
		num = -num
	}

	var count int
	// Repeatedly divide by 10 until the number becomes 0
	for num > 0 {
		num = num / 10
		count++
	}
	return count
}

// countDigitsRecursive counts the total number of digits in an integer using recursion.
//
// Time Complexity: O(log10(n)), Space Complexity: O(log10(n)) (call stack)
func countDigitsRecursive(num int) int {
	// Handle negative numbers
	if num < 0 {
		num = -num
	}

	// Base case: single-digit number (0 to 9)
	if num < 10 {
		return 1
	}

	// Recursive step: 1 + count of remaining digits
	return 1 + countDigitsRecursive(num/10)
}
