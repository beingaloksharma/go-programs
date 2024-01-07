// // WAP
// // To print even and odd number using goroutines and channels
// package main

// import "fmt"

// func main() {
// 	evenCh := make(chan int)
// 	oddCh := make(chan int)
// 	go evenNumber(evenCh)
// 	go oddNumber(oddCh)
// 	for {
// 		select {
// 		case enum, ok := <-evenCh:
// 			if ok {
// 				fmt.Println("Even Number is :: ", enum)
// 			} else {
// 				evenCh = nil
// 			}
// 		case onum, ok := <-oddCh:
// 			if ok {
// 				fmt.Println("odd Number is :: ", onum)
// 			} else {
// 				oddCh = nil
// 			}
// 		}
// 		if evenCh == nil && oddCh == nil {
// 			break
// 		}
// 	}
// }

// func evenNumber(even chan<- int) {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			even <- i
// 		}
// 	}
// 	close(even)
// }

// func oddNumber(odd chan<- int) {
// 	for i := 0; i <= 10; i++ {
// 		if i%2 != 0 {
// 			odd <- i
// 		}
// 	}
// 	close(odd)
// }
