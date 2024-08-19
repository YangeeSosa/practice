package kata

import (
	"math"
)

func FindNextSquare(sq int64) int64 {
	n := math.Sqrt(float64(sq))
	if math.Mod(n, 1.00) == 0 {
		nxt := (n + 1) * (n + 1)
		return int64(nxt)
	}
	return -1
}
