// Package main provides a basic program to find the largest and smallest numbers in a slice of integers in Go.
//
// Program: Find Largest and Smallest Number
// Description: Iterates through a slice of integers to find the minimum and maximum values
// while handling edge cases such as empty slices.
package main

import "fmt"

func main() {
	// Sample slice of integers to evaluate
	nums := []int{15, 12, 9, 15, 17, 4, 16, 18}

	fmt.Println("Numbers          ::", nums)
	fmt.Println("Smallest Number  ::", findSmallest(nums))
	fmt.Println("Largest Number   ::", findLargest(nums))
}

// findSmallest finds the minimum value in a slice of integers.
//
// Time Complexity: O(n), Space Complexity: O(1)
func findSmallest(nums []int) int {
	// Handle the edge case of an empty slice to prevent panics
	if len(nums) == 0 {
		return 0
	}

	// Assume the first number is the smallest initially
	smallest := nums[0]

	// Iterate through the rest of the slice
	for i := 1; i < len(nums); i++ {
		// If we find a number smaller than our current smallest, update it
		if nums[i] < smallest {
			smallest = nums[i]
		}
	}

	return smallest
}

// findLargest finds the maximum value in a slice of integers.
//
// Time Complexity: O(n), Space Complexity: O(1)
func findLargest(nums []int) int {
	// Handle the edge case of an empty slice to prevent panics
	if len(nums) == 0 {
		return 0
	}

	// Assume the first number is the largest initially
	largest := nums[0]

	// Iterate through the rest of the slice
	for i := 1; i < len(nums); i++ {
		// If we find a number greater than our current largest, update it
		if nums[i] > largest {
			largest = nums[i]
		}
	}

	return largest
}
