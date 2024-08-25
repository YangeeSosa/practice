package kata

import (
	"strings"
)

func SizeToNumber(size string) (int, bool) {
	switch {
	case size == "s":
		return 36, true
	case size == "m":
		return 38, true
	case size == "l":
		return 40, true
	case size == "":
		return 0, false
	}
	amountOfX := strings.Count(size, "x")
	if size[amountOfX:] == "l" {
		return 40 + (amountOfX * 2), true
	} else if size[amountOfX:] == "s" {
		return 36 - (amountOfX * 2), true
	}
	return 0, false
}
