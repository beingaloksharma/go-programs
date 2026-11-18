// Package main provides a basic program to check if a number is an Armstrong number in Go.
//
// Program: Armstrong Number Check
// Description: Determines whether a given number is an Armstrong number (narcissistic number),
// where the sum of its own digits each raised to the power of the number of digits
// equals the number itself.
package main

import "fmt"

func main() {
	// Sample integer to test
	num := 153
	copyNum := num
	count := 0

	// Count the number of digits and demonstrate digit extraction
	for num != 0 {
		lastDigit := num % 10
		num = num / 10
		count = count + 1
		fmt.Println("Extracting digit:", lastDigit)
	}

	// Verify and print if the number is an Armstrong number
	if armstrongNumber(copyNum, count) {
		fmt.Printf("%d is an Armstrong Number \n", copyNum)
	} else {
		fmt.Printf("%d is not an Armstrong Number \n", copyNum)
	}
}

// armstrongNumber checks if a number equals the sum of its digits raised to the power of count.
//
// Time Complexity: O(d * count) where d is the number of digits, Space Complexity: O(1)
func armstrongNumber(num, count int) bool {
	copyNum := num
	var total int

	// Extract each digit, raise to power of count, and accumulate sum
	for num > 0 {
		digit := num % 10
		total += powerOfNum(digit, count)
		num = num / 10
	}

	// Check if the accumulated sum equals the original number
	return total == copyNum
}

// powerOfNum calculates the result of num raised to the given power (num^pow).
//
// Time Complexity: O(pow), Space Complexity: O(1)
func powerOfNum(num, pow int) int {
	res := 1
	for i := 0; i < pow; i++ {
		res *= num
	}
	return res
}
