package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println(<-ch)
	}(ch, wg)
	go func(ch chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- "Sending message to receiver channel"
	}(ch, wg)
	wg.Wait()
}
