// Package main provides a basic program to check whether a number is even or odd in Go.
//
// Program: Even or Odd Number Check
// Description: Determines whether a given integer is even or odd using the modulo operator.
package main

import "fmt"

func main() {
	// Sample integers to test for even or odd
	nums := []int{11, 13, 15, 17, 19, 20, 22, 24, 26, 28, 32, 43, 45}

	// Iterate over test numbers and check if each is even or odd
	for _, num := range nums {
		if isNumEvenOdd(num) {
			fmt.Printf("%d is even number \n", num)
		} else {
			fmt.Printf("%d is odd number \n", num)
		}
	}
}

// isNumEvenOdd checks whether a given integer is even.
// Returns true if the number is even (divisible by 2 with no remainder), false if it is odd.
//
// Time Complexity: O(1), Space Complexity: O(1)
func isNumEvenOdd(num int) bool {
	return num%2 == 0
}
