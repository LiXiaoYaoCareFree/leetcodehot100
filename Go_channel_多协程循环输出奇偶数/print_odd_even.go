package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	oddCh, evenCh := make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(2)

	go printOdd(oddCh, evenCh, &wg)
	go printEven(oddCh, evenCh, &wg)

	oddCh <- struct{}{}

	wg.Wait()
	close(oddCh)
	close(evenCh)
}

func printEven(oddCh, evenCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	num := 2
	for {
		<-evenCh
		fmt.Printf("%d ", num)
		time.Sleep(500 * time.Millisecond)
		num += 2
		oddCh <- struct{}{}
	}
}

func printOdd(oddCh, evenCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	num := 1
	for {
		<-oddCh
		fmt.Printf("%d ", num)
		time.Sleep(500 * time.Millisecond)
		num += 2
		evenCh <- struct{}{}
	}
}
