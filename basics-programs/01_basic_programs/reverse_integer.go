// Package main provides a basic program to reverse an integer and check for palindrome in Go.
//
// Program: Reverse Integer & Palindrome Check
// Description: Reverses the digits of an integer using arithmetic operations
// and determines whether it is a palindrome.
package main

import "fmt"

func main() {
	// Sample integers to test
	nums := []int{12345, 54321, 1234321, 5678765}

	// Iterate over test numbers, reverse their digits, and check palindrome property
	for _, num := range nums {
		revOfNum := reverseOfNumber(num)
		if revOfNum == num {
			fmt.Printf("%d is a palindrome number\n", num)
		} else {
			fmt.Printf("%d is not a palindrome number\n", num)
		}
	}
}

// reverseOfNumber reverses the digits of an integer using modulo and division.
//
// Time Complexity: O(log10(n)), Space Complexity: O(1)
func reverseOfNumber(num int) int {
	// Variable to store the reversed number
	rev := 0

	// Extract digits from right to left and append to rev
	for num > 0 {
		r := num % 10
		rev = rev*10 + r
		num = num / 10
	}

	// Return reversed number
	return rev
}
