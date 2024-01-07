// package main

// import "fmt"

// func main() {
// 	msg := make(chan string) // unbuffered channel
// 	go func() {
// 		msg <- "Hello World!. This is example of Goroutines and Channels"
// 	}()
// 	message := <-msg
// 	fmt.Println("Channel is receving Messgae :: ", message)
// }
