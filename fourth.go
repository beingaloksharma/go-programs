// package main

// import "fmt"

// func main() {
// 	ch := make(chan int)
// 	go checkNumber(ch)
// 	ch <- 11
// }

// func checkNumber(ch chan int) {
// 	if <-ch%2 == 0 {
// 		fmt.Println("I am even number")
// 	} else {
// 		fmt.Println("I am odd number")
// 	}
// }
