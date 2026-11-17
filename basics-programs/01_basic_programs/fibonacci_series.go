// Package main provides a basic program to generate the Fibonacci series in Go.
//
// Program: Fibonacci Series
// Description: Generates the first N numbers of the Fibonacci sequence using
// an iterative calculation with case handling via switch statement.
package main

import "fmt"

func main() {
	// Sample slice representing the count of Fibonacci numbers to generate
	nums := []int{0, 5, 8, 2, 6, 10, 3, 7, 4, 1, 9}

	// Iterate over each test value and generate its Fibonacci sequence
	for _, num := range nums {
		// Initialize the first two terms of the Fibonacci sequence
		start := 0
		next := 1

		fmt.Printf("--- Generating Fibonacci for %d ---\n", num)

		switch {
		// Handle invalid input (less than or equal to 0)
		case num <= 0:
			fmt.Println("Please enter a valid number (greater than 0)")

		// Base case: first term
		case num == 1:
			fmt.Println("Fibonacci Series Number ::", start)

		// Base case: first two terms
		case num == 2:
			fmt.Println("Fibonacci Series Number ::", start)
			fmt.Println("Fibonacci Series Number ::", next)

		// For 3 or more terms, iteratively calculate and display the sequence
		case num >= 3:
			fmt.Println("Fibonacci Series Number ::", start)
			fmt.Println("Fibonacci Series Number ::", next)

			for i := 3; i <= num; i++ {
				res := fibonacciSeries(start, next)
				start = next
				next = res
				fmt.Println("Fibonacci Series Number ::", res)
			}
		}
		fmt.Print("-------------------------\n\n")
	}
}

// fibonacciSeries calculates the next number in the Fibonacci sequence
// by summing the two preceding numbers.
//
// Time Complexity: O(1), Space Complexity: O(1)
func fibonacciSeries(start, next int) int {
	return start + next
}
