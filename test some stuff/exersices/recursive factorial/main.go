package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Высисляем...\n")

	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	fmt.Println(factorial(45))

}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	time.Sleep(time.Millisecond * 3)
	return n * factorial(n-1)
}
