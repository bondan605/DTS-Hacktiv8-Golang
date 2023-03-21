package main

import (
	"fmt"
	"sync"
)

func goroutineCoba(wg *sync.WaitGroup, mu *sync.RWMutex, i int, data []string, ch1 *chan string) {
	defer wg.Done()
	*ch1 <- fmt.Sprintf("%v %d", data, i)
}

func goroutineBisa(wg *sync.WaitGroup, mu *sync.RWMutex, i int, data []string, ch2 *chan string) {
	defer wg.Done()
	*ch2 <- fmt.Sprintf("%v %d", data, i)
}

func goroutineCobaBisa(wg *sync.WaitGroup, mu *sync.RWMutex, ch1 *chan string, ch2 *chan string, chJoin *chan string) {
	defer wg.Done()
	mu.Lock()
	*chJoin <- <-*ch1
	*chJoin <- <-*ch2
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	var chJoin = make(chan string, 2)
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	var data1 = []string{"coba1", "coba2", "coba3"}
	var data2 = []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go goroutineCoba(&wg, &mu, i, data1, &ch1)
		go goroutineBisa(&wg, &mu, i, data2, &ch2)
	}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go goroutineCobaBisa(&wg, &mu, &ch1, &ch2, &chJoin)
	}

	for i := 1; i <= 8; i++ {
		fmt.Println(<-chJoin)
	}

	go func() {
		wg.Wait()
		close(chJoin)
	}()
}
