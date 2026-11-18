// Package main provides basic programs to swap two numbers in Go.
//
// Program: Swap Two Numbers
// Methods demonstrated:
//  1. Idiomatic Go multiple assignment (tuple assignment)
//  2. Arithmetic operations (without a temporary variable)
//  3. Bitwise XOR operator (without a temporary variable)
package main

import "fmt"

func main() {
	num1, num2 := 10, 20

	fmt.Println("Original Values :: Number 1 =", num1, ", Number 2 =", num2)

	// Approach 1: Idiomatic Go tuple assignment
	fmt.Println("\n--- Approach 1: Tuple Assignment ---")
	swapNums(num1, num2)

	// Approach 2: Arithmetic operations (without temp variable)
	fmt.Println("\n--- Approach 2: Arithmetic Swap ---")
	swapArithmetic(num1, num2)

	// Approach 3: Bitwise XOR (without temp variable)
	fmt.Println("\n--- Approach 3: Bitwise XOR Swap ---")
	swapXOR(num1, num2)
}

// swapNums swaps two numbers using Go's idiomatic multiple assignment.
//
// Time Complexity: O(1), Space Complexity: O(1)
func swapNums(num1, num2 int) {
	num1, num2 = num2, num1
	fmt.Println("Swap Value : Number 1 :: ", num1)
	fmt.Println("Swap Value : Number 2 :: ", num2)
}

// swapArithmetic swaps two numbers using addition and subtraction without a temporary variable.
//
// Time Complexity: O(1), Space Complexity: O(1)
func swapArithmetic(num1, num2 int) {
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println("Swap Value : Number 1 :: ", num1)
	fmt.Println("Swap Value : Number 2 :: ", num2)
}

// swapXOR swaps two numbers using bitwise XOR without a temporary variable.
//
// Time Complexity: O(1), Space Complexity: O(1)
func swapXOR(num1, num2 int) {
	num1 = num1 ^ num2
	num2 = num1 ^ num2
	num1 = num1 ^ num2
	fmt.Println("Swap Value : Number 1 :: ", num1)
	fmt.Println("Swap Value : Number 2 :: ", num2)
}
