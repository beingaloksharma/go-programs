// Package main provides a basic program to check if a string is a palindrome in Go.
//
// Program: String Palindrome Check
// Description: Checks whether a given string reads the same forwards and backwards.
package main

import "fmt"

func main() {
	// Sample strings to test
	strings := []string{"madam", "golang", "mom"}

	// Iterate over test strings and check each for palindrome property
	for _, val := range strings {
		if isPalindrome(val) {
			fmt.Printf("%s is palindrome\n", val)
		} else {
			fmt.Printf("%s is not palindrome\n", val)
		}
	}
}

// isPalindrome checks whether a given string is a palindrome by reversing
// its runes and comparing the reversed string with the original.
//
// Time Complexity: O(n), Space Complexity: O(n)
func isPalindrome(str string) bool {
	// Convert string to rune slice to handle multi-byte Unicode characters properly
	chars := []rune(str)

	// Reverse the rune slice in-place using two pointers
	// Note: len(chars) is used instead of len(str) to avoid out-of-range panics on multi-byte runes
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	// Compare original string with reversed string and return result
	return str == string(chars)
}
