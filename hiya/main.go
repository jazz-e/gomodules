package main

import (
	"fmt"
	"time"
)

func f() {
	x := "Hello Code Nijas!"

	fmt.Println(x)
}

func main() {
	go f()

	time.Sleep(1 * time.Second)

	fmt.Println(("main function"))

}
