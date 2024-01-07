// package main

// import "fmt"

// func main() {
// 	evenCh := make(chan int)
// 	oddCh := make(chan int)
// 	go evenNumbers(evenCh)
// 	go oddNumbers(oddCh)
// 	for {
// 		select {
// 		case num, ok := <-evenCh:
// 			if ok {
// 				fmt.Println("Even Number is :: ", num)
// 			} else {
// 				evenCh = nil
// 			}
// 		case num, ok := <-oddCh:
// 			if ok {
// 				fmt.Println("Odd Number is :: ", num)
// 			} else {
// 				oddCh = nil
// 			}
// 		}
// 		if evenCh == nil && oddCh == nil {
// 			break
// 		}
// 	}
// }

// func evenNumbers(evenCh chan int) {
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			evenCh <- i
// 		}
// 	}
// 	close(evenCh)
// }

// func oddNumbers(oddCh chan int) {
// 	for i := 0; i < 10; i++ {
// 		if i%2 != 0 {
// 			oddCh <- i
// 		}
// 	}
// 	close(oddCh)
// }
