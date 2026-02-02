package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	catCh, dogCh, fishCh := make(chan struct{}), make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)

	go printCat(catCh, dogCh, fishCh, &wg)
	go printDog(catCh, dogCh, fishCh, &wg)
	go printFish(catCh, dogCh, fishCh, &wg)

	catCh <- struct{}{}

	time.Sleep(10 * time.Millisecond)

	wg.Wait()
	close(catCh)
	close(dogCh)
	close(fishCh)
}

func printFish(catCh, dogCh, fishCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-fishCh:
			fmt.Print("fish ")
			time.Sleep(time.Millisecond * 500)
			catCh <- struct{}{}
		}
	}
}

func printDog(catCh, dogCh, fishCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-dogCh:
			fmt.Print("dog ")
			time.Sleep(time.Millisecond * 500)
			fishCh <- struct{}{}
		}
	}
}

func printCat(catCh, dogCh, fishCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-catCh:
			fmt.Print("cat ")
			time.Sleep(time.Millisecond * 500)
			dogCh <- struct{}{}
		}
	}
}
