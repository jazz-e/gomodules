package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1) // increment to a counter
		go func(j int) {
			defer wg.Done() // decrement the counter
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)

			if err != nil {
				// Port is closed or filtered
				return
			}

			conn.Close()

			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait() //wait till counter gets to nil
}
