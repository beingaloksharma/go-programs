//Write a program to reverse a string

// package main

// import "fmt"

// func main() {
// 	str := "Hello World !!!"
// 	revStr := reverseString(str)
// 	fmt.Println("String is :: ", str)
// 	fmt.Println("Reverse of String :: ", revStr)
// }

// func reverseString(str string) string {
// 	runes := []rune(str)
// 	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }
