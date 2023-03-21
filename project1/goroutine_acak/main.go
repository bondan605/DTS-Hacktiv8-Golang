package main

import (
	"fmt"
	"sync"
)

func goroutineTest(wg *sync.WaitGroup, mu *sync.Mutex, i int, data []string, ch chan string) {
	defer wg.Done()
	ch <- fmt.Sprintf("%v %d", data, i)
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var ch = make(chan string)
	var data1 = []string{"coba1", "coba2", "coba3"}
	var data2 = []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go goroutineTest(&wg, &mu, i, data1, ch)
		go goroutineTest(&wg, &mu, i, data2, ch)
	}

	for i := 1; i <= 8; i++ {
		fmt.Println(<-ch)
	}

	wg.Wait()
}
