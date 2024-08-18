package kata

import (
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {
	parts := strings.Split(time, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])

	hourIn12 := hour % 12
	if hourIn12 == 0 {
		hourIn12 = 12
	}

	switch {
	case minute == 0:
		return strings.Repeat("Cuckoo ", hourIn12)[:len(strings.Repeat("Cuckoo ", hourIn12))-1]
	case minute == 30:
		return "Cuckoo"
	case minute%15 == 0:
		return "Fizz Buzz"
	case minute%5 == 0:
		return "Buzz"
	case minute%3 == 0:
		return "Fizz"
	default:
		return "tick"
	}
}
