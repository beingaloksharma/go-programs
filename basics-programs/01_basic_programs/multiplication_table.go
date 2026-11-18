// Package main provides a basic program to generate the multiplication table of a number in Go.
//
// Program: Multiplication Table
// Description: Prints the multiplication table for a given integer from 1 to 10
// using formatted output.
package main

import "fmt"

func main() {
	// Sample integer to generate multiplication table for
	table := 5
	fmt.Printf("--- Multiplication Table for %d ---\n", table)
	multiplicationTable(table)
}

// multiplicationTable prints the multiplication table of a given number from 1 to 10.
//
// Time Complexity: O(1) (fixed 10 iterations), Space Complexity: O(1)
func multiplicationTable(num int) {
	// Iterate from 1 to 10 and print each multiplication expression
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
	fmt.Println()
}
