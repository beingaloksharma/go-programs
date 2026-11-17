// Package main provides basic programs to reverse a string in Go.
//
// Program: Reverse a string
// Methods demonstrated:
//  1. Backward iteration using rune slice
//  2. In-place two-pointer swap using rune slice
package main

import "fmt"

func main() {
	// Test strings (ASCII and Unicode)
	str1 := "hello world!"
	str2 := "Hello, 世界! 🚀"

	fmt.Println("--- Test Case 1 (ASCII) ---")
	fmt.Println("Original String   ::", str1)
	fmt.Println("Approach 1 Result ::", reverseString_one(str1))
	fmt.Println("Approach 2 Result ::", reverseString_two(str1))

	fmt.Println("\n--- Test Case 2 (Unicode & Emojis) ---")
	fmt.Println("Original String   ::", str2)
	fmt.Println("Approach 1 Result ::", reverseString_one(str2))
	fmt.Println("Approach 2 Result ::", reverseString_two(str2))
}

// reverseString_one reverses a string by iterating backwards over its runes.
// Time Complexity: O(n), Space Complexity: O(n)
func reverseString_one(str string) string {
	// Convert string to rune slice to handle multi-byte Unicode characters properly
	runes := []rune(str)

	// Check length of string
	if len(runes) == 0 {
		return ""
	}

	// Pre-allocate slice with known capacity for performance
	chars := make([]rune, 0, len(runes))

	// Iterate backwards over runes
	for i := len(runes) - 1; i >= 0; i-- {
		chars = append(chars, runes[i])
	}

	// Return result as string
	return string(chars)
}

// reverseString_two reverses a string in-place using the two-pointer technique.
// Time Complexity: O(n), Space Complexity: O(n)
func reverseString_two(str string) string {
	// Convert string to rune slice for in-place swapping
	chars := []rune(str)

	// Two-pointer swap
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	// Return result as string
	return string(chars)
}
