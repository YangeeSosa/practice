package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := make(chan int)

	go func() {
		numbers <- 42
	}()

	time.Sleep(time.Millisecond * 100)

	select {
	case n := <-numbers:
		fmt.Println(n)
	default:
		fmt.Println("Ничего, пусто, complete nothing")
	}
}
