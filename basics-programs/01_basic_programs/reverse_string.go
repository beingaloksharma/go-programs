package main

import "fmt"

func main() {
	//To declare a variable
	str := "hello world!"
	//To store result
	result_one := reverseString_one(str)
	//Print result
	fmt.Println("Reverse of String :: ", result_one)
}

func reverseString_one(str string) string {
	//To store result
	var chars []rune
	//Check lenght of string
	if len(str) == 0 {
		return ""
	}
	//Iterate over string
	for i := len(str) - 1; i >= 0; i-- {
		chars = append(chars, rune(str[i]))
	}
	//Return result
	return string(chars)
}
