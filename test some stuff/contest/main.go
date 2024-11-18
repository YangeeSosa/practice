package main

import (
	"fmt"
)

var (
	queue      int
	cupBus     int
	middleSeat int
	lastCupBus int
)

func main() {
	fmt.Scan(&queue, &cupBus)
	if cupBus%2 == 0 {
		middleSeat = cupBus / 2
	} else {
		middleSeat = cupBus/2 + 1
	}
	for queue > 0 {
		nextSeat := middleSeat
		for i := 1; lastCupBus < cupBus; i++ {
			switch {
			case cupBus%2 == 0:
				if i%2 != 0 {
					nextSeat = nextSeat - i + 1
				} else {
					nextSeat = nextSeat + i - 1
				}
			default:
				if i%2 != 0 {
					nextSeat = nextSeat + i - 1
				} else {
					nextSeat = nextSeat - i + 1
				}
			}
			fmt.Println(nextSeat)
			queue--
			lastCupBus++
			if queue == 0 {
				break
			}
		}
		lastCupBus = 0
	}
}
