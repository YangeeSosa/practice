package leetcode

import "math"

func reverse(x int) int {
	original := x
	reversed := 0

	for original != 0 {
		reversed = reversed*10 + original%10
		original /= 10
	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	return reversed
}
