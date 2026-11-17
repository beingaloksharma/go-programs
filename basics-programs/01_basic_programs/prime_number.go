// Package main provides a basic program to check for prime numbers in Go.
//
// Program: Prime Number Check
// Description: Determines whether a given number is prime by checking divisibility
// up to its square root, and counts primes in a given range.
package main

import "fmt"

func main() {
	// Counter for total prime numbers in the range
	var numOfPrime int

	// Sample integers to test
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Iterate over test numbers and check if each is a prime number
	for _, num := range nums {
		if IsPrime(num) {
			fmt.Printf("%d is a prime number \n", num)
		} else {
			fmt.Printf("%d is not a prime number \n", num)
		}
	}

	// Count and display all prime numbers in the range 0 to 99
	for i := 0; i < 100; i++ {
		if IsPrime(i) {
			fmt.Printf("%d is a prime number \n", i)
			numOfPrime += 1
		} else {
			fmt.Printf("%d is not a prime number \n", i)
		}
	}
	fmt.Println("Total Prime from 0 to 100 are :: ", numOfPrime)
}

// IsPrime checks whether a given integer is a prime number.
// It optimizes the check by testing divisibility only up to the square root of the number.
//
// Time Complexity: O(sqrt(n)), Space Complexity: O(1)
func IsPrime(num int) bool {
	// Numbers less than 2 are not prime numbers
	if num < 2 {
		return false
	}

	// Check divisibility from 2 up to sqrt(num)
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	// Return true if no divisors were found
	return true
}
