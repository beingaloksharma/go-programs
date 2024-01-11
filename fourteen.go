// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	str := "ABA"
// 	res := reverse(str)
// 	if reflect.DeepEqual(str, res) {
// 		fmt.Println("String is palindrome")
// 	} else {
// 		fmt.Println("String is not palindrome")
// 	}
// }

// func reverse(str string) string {
// 	res := []rune(str)
// 	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
// 		res[i], res[j] = res[j], res[i]
// 	}
// 	return string(res)
// }
