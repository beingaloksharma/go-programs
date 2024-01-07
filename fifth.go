// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch := make(chan int)
// 	go Table(ch)
// 	ch <- 10
// }

// func Table(ch chan int) {
// 	fmt.Printf("Table of %d \n", <-ch)
// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("%d * %d = %d \n", ch, i, <-ch*i)
// 	}
// }
