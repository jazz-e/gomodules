package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)      // create a channel and buffer size
	var wg sync.WaitGroup             // create a wait group structure
	for i := 0; i < cap(ports); i++ { // start each worker x 100
		go worker(ports, &wg) // coroutines. ports (100) and wait group
	}
	for i := 1; i < 1024; i++ { // 1024 ports
		wg.Add(1) // increment wait group
		ports <- i
	}
	wg.Wait()
	close(ports)
}
